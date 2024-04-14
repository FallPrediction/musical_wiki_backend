package utils

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Cache struct {
	RedisClient *redis.Client
	logger      *zap.SugaredLogger
}

func (c *Cache) Get(key string, v any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cacheBytes, err := c.RedisClient.Get(ctx, key).Bytes()
	if err == nil {
		err = json.Unmarshal(cacheBytes, v)
		if err != nil {
			c.logger.Warn("json unmarshal error", err)
		}
	}
	return err
}

func (c *Cache) Set(key string, v any) error {
	cacheBytes, err := json.Marshal(v)
	if err != nil {
		c.logger.Warn("json marshal error", err)
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()
		err = c.RedisClient.Set(ctx, key, cacheBytes, 24*time.Hour).Err()
	}
	return err
}

func (c *Cache) Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := c.RedisClient.Del(ctx, key).Err()
	if ctx.Err() == context.DeadlineExceeded {
		c.logger.Warn("Delete cache timeout, key: ", key)
	}
	return err
}

func (c *Cache) ScanAndDel(match string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	iter := c.RedisClient.Scan(ctx, 0, match, 0).Iterator()
	for iter.Next(ctx) {
		err := c.Del(iter.Val())
		if err != nil {
			c.logger.Warn("Delete cache error, key: ", iter.Val())
		}
	}
	if ctx.Err() == context.DeadlineExceeded {
		c.logger.Warn("Scan and delete cache timeout, match: ", match)
	}
}

func NewCache(redis *redis.Client, logger *zap.SugaredLogger) Cache {
	return Cache{
		RedisClient: redis,
		logger:      logger,
	}
}
