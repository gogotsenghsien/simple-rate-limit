package caches

import (
	"github.com/go-redis/redis"
	"github.com/gogotsenghsien/simple-rate-limit/src/configs"
)

type Redis struct {
	*redis.Client
}

func NewRedis(config *configs.Config) (*Redis, error) {
	host := config.GetString("redis.url")
	password := config.GetString("redis.password")
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return &Redis{client}, nil
}
