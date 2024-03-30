package initialize

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var redisClient *redis.Client
var redisOnce sync.Once

func NewRedis(logger *zap.SugaredLogger) *redis.Client {
	if redisClient == nil {
		redisOnce.Do(func() {
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
				logger.Error("redis connect ping failed, err: ", err)
			}
			redisClient = client
		})
	}
	return redisClient
}
