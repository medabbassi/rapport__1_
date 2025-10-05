package cache

import (
	"context"
	"github.com/medabbassi/go_server/pkg/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var cacheClient *redis.Client
var ctx = context.Background()

func InitCache(cfg *config.Config) {
	cacheClient = redis.NewClient(&redis.Options{
		Addr: cfg.REDIS, // Use the same config field for Redis address
	})
	_, err := cacheClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis connection error: %v", err)
	}
}

func CacheSet(key, value string, ttl int32) {
	err := cacheClient.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		log.Printf("Cache set error: %v", err)
	}
}

func CacheGet(key string) (string, error) {
	val, err := cacheClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	}
	if err != nil {
		return "", err
	}
	return val, nil
}
