package dbTool

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func init() {
	eventMap = make(map[string]func(interface{}))
}

var eventMap map[string]func(interface{})

func register(name string, f func(interface{})) {
	eventMap[name] = f
}

//处理延时处理数据
func handle(stru interface{}, key string) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	var val *redis.StringSliceCmd
	RedisClient.Pipelined(context.Background(), func(pipeliner redis.Pipeliner) error {
		val = pipeliner.ZRangeByScore(context.Background(), key, &redis.ZRangeBy{
			Min: "0",
			Max: now,
		})
		pipeliner.ZRemRangeByScore(context.Background(), key, "0", now)
		return nil
	})
	strs, err := val.Result()
	if err != nil {
		panic(err)
	}
	for _, str := range strs {
		json.Unmarshal([]byte(str), stru)
		go eventMap[key](stru)
	}
}
