package timer

import (
	"context"
	"time"
)

func Register(t int64, f func()) {
	var cancel context.CancelFunc
	con := context.Background()
	con, cancel = context.WithTimeout(con, time.Duration(t)*time.Second)
	con = context.WithValue(con, "time", t)
	go handler(con, f, cancel)
}

func handler(con context.Context, f func(), cancel context.CancelFunc) {
	select {
	case <-con.Done():
		go Register(con.Value("time").(int64), f)
		go f()
		cancel()
	}
}
