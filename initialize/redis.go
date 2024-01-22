package initialize

import (
	"context"
	"fmt"
	"musical_wiki/global"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:6379", os.Getenv("REDIS_ADDR")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("redis connect ping failed, err: ", err)
	} else {
		global.Redis = client
	}
}
