package dbTool

import (
	"context"
	"encoding/json"
	"github.com/chen-mou/gf/os/glog"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/http2"
	"golang.org/x/sys/windows"
	"os"
	"strconv"
	"time"
)

var RedisClient *redis.ClusterClient

var redisPath = []string{
	"150.158.169.43:6371",
	"150.158.169.43:6372",
	"150.158.169.43:6373",
}

var localRedisPath = []string{
	"127.0.0.1:6380",
	"127.0.0.1:6381",
	"127.0.0.1:6382",
}

var myRedisPath = []string{
	"120.24.214.131:6381",
	"120.24.214.131:6382",
	"120.24.214.131:6383",
}

func do(handler func(), ctx context.Context) {
	select {
	case <-ctx.Done():
		handler()
	}
}

func init() {
	RedisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    redisPath,
		Password: "1007324849redis...",
	})
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ok := RedisClient.Ping(ctx)
	if ok.Val() == "PONG" {
		glog.Info("Redis init success")
	} else {
		panic(any("Redis init error err is : " + ok.Err().Error()))
	}
	//ctx1, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//glog.Info(redis.NewClient(&redis.Options{
	//	Addr:     "150.158.169.43:6371",
	//	Password: "1007324849redis...",
	//}).Ping(ctx1).String())
}

// GetLock 获取锁
// 返回 true 或者 false;
func GetLock(key, value string, expireAt time.Duration) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cmd bool
	var ok error
	cmd, ok = RedisClient.SetNX(ctx, key, value, expireAt).Result()
	if ok != nil {
		panic(any(ok))
	}
	return cmd
}

// GetLoopLock 自旋锁直到获取锁或者超时否则一直阻塞
// deadLine 取-1为永不超时
func GetLoopLock(key, value string, expireAt time.Duration, deadLine int64) bool {
	ctx, cancelFunc1 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc1()
	var cmd bool
	var ok error
	var timeout = false
	if deadLine != -1 {
		ctx1, cancelFunc := context.WithTimeout(context.Background(),
			time.Duration(deadLine)*time.Millisecond)
		defer cancelFunc()
		go do(func() {
			timeout = !timeout
		}, ctx1)
	}
	for !cmd && !timeout {
		cmd, ok = RedisClient.SetNX(ctx, key, value, expireAt).Result()
		if ok != nil {
			panic(any(ok))
		}
		time.Sleep(100 * time.Millisecond)
	}
	return cmd
}

func Unlock(key, value string) string {
	ctx, f := context.WithTimeout(context.Background(), 10*time.Second)
	defer f()
	script := redis.NewScript(`
	local key = KEYS[1]
	local value = ARGV[1]
	local back = redis.call("GET", key)
	if not value then
		return "unknown"
	end
	if value == back then
		redis.call("DEL", key)
		return "ok"
	end
	return "wait"`)
	v, err := script.Run(ctx, RedisClient, []string{key}, []string{value}).Result()
	if err != nil {
		panic(any(err))
	}
	return v.(string)

}

func Get(key string, value interface{}) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	str, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		panic(any(err))
	}
	json.Unmarshal([]byte(str), value)
}

func Set(key string, value interface{}, expireAt time.Duration) {
	str, err := json.Marshal(value)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	if err != nil {
		panic(any(err))
	}
	RedisClient.Set(ctx, key, str, expireAt)
}

func GetThreadID() string {
	goroutineID := http2.CurGoroutineID()
	pid := os.Getpid()
	threadId := windows.GetCurrentThreadId()
	return strconv.Itoa(pid) + ":" +
		strconv.FormatUint(uint64(threadId), 10) + ":" +
		strconv.FormatUint(goroutineID, 10)
}
