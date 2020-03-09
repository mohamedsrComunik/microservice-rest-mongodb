package cache

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	client *redis.Client
}

func NewRedis() redisCache{
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}
	)
	return redisCache{client}
}

func (self redisCache) Get(key string){
	result,err := self.client.Get(key).Result()
	if err == nil {
		return result 
	}
	return ""
	
}

func (self redisCache) Set(key string, value string, expire time.Duration) error {
	return self.Client.Set(key, value, expire).Err()
}