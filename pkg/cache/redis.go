package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (redis *RedisCache) Get(key string) (value string) {
	return redis.client.Get(key).Val()
}

func (redis *RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	return redis.client.Set(key, value, ttl).Err()
}
