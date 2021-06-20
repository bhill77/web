package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisCacheRepository struct {
	client *redis.Client
}

// NewRedisCacheRepository new redis cache service
func NewRedisCacheRepository(client *redis.Client) *RedisCacheRepository {
	return &RedisCacheRepository{
		client: client,
	}
}

// Set set cache based on given key
func (redisCache *RedisCacheRepository) Set(key string, value string, expiration time.Duration) error {

	command := redisCache.client.Set(key, value, expiration)

	if command.Err() != nil {
		return command.Err()
	}

	return nil
}
