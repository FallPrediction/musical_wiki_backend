package initialize

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewRedis(logger *zap.SugaredLogger) *redis.Client {
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
	} else {
		return client
	}
	return nil
}
