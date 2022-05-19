package client

import (
	"context"
	"google.golang.org/grpc"
	"sync"
	"time"
)

type Connect struct {
	*grpc.ClientConn
	con    context.Context
	index  int
	isDone bool
}

type connPool struct {
	pool     []*Connect
	maxNum   int
	coreNum  int
	overtime time.Duration
	queue    queue
}

type queue struct {
	sync.Mutex
	que   []chan *Connect
	first int
	end   int
}

var pool connPool

var url = "localhost:14000"

func (q *queue) push(con chan *Connect) bool {
	q.Lock()
	defer q.Unlock()
	if (q.end+1)%cap(q.que) == q.first {
		return false
	}
	q.que[q.end] = con
	q.end = (q.end + 1) % cap(q.que)
	return true
}

func (q *queue) pop() (chan *Connect, bool) {
	q.Lock()
	defer q.Unlock()
	if q.end == q.first {
		return nil, false
	}
	val := q.que[q.first]
	q.first = (q.first + 1) % cap(q.que)
	return val, true
}

func (c *Connect) Done() {
	ch, ok := pool.queue.pop()
	if !ok {
		c.isDone = true
		c.con, _ = context.WithTimeout(context.Background(), 5*time.Minute)
		go overdue(c)
		return
	}
	ch <- c
}

func init() {
	pool = connPool{
		pool:     make([]*Connect, 8),
		maxNum:   8,
		coreNum:  6,
		overtime: 10 * time.Minute,
		queue: queue{
			que:   make([]chan *Connect, 20),
			first: 0,
			end:   0,
		},
	}
}

var poolLock sync.Mutex

// overdue 处理存活时间超时的连接
func overdue(connect *Connect) {
	select {
	case <-connect.con.Done():
		poolLock.Lock()
		pool.pool = append(pool.pool[:connect.index], pool.pool[connect.index+1:]...)
		connect.Close()
		poolLock.Unlock()
	}
}

func handler(ch chan *Connect) bool {
	if len(pool.pool) < pool.coreNum {
		poolLock.Lock()
		conn, err := grpc.Dial(url)
		if err != nil {
			panic(err)
		}
		val := &Connect{
			isDone:     false,
			ClientConn: conn,
			index:      len(pool.pool),
		}
		ch <- val
		pool.pool = append(pool.pool, val)
		poolLock.Unlock()
		return true
	} else {
		if pool.queue.push(ch) {
			return true
		} else {
			if cap(pool.pool) > len(pool.pool) {
				poolLock.Lock()
				conn, err := grpc.Dial(url)
				if err != nil {
					panic(err)
				}
				val := &Connect{
					isDone:     false,
					ClientConn: conn,
					index:      len(pool.pool),
				}
				ch <- val
				pool.pool = append(pool.pool, val)
				poolLock.Unlock()
				return true
			}
			return false
		}
	}
}

func getConnect() (<-chan *Connect, bool) {
	ch := make(chan *Connect, 1)
	ok := handler(ch)
	return ch, ok
}

func GetConnect(ctx context.Context) (*Connect, bool) {
	ch, ok := getConnect()
	select {
	case val := <-ch:
		return val, ok
	case <-ctx.Done():
		return nil, false
	}
}
