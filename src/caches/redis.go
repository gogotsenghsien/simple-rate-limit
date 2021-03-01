package caches

import (
	"github.com/go-redis/redis"
	"github.com/gogotsenghsien/simple-rate-limit/src/configs"
	"github.com/gogotsenghsien/simple-rate-limit/src/logs"
)

type Redis struct {
	*redis.Client
}

func NewRedis(config *configs.Config, logger *logs.Logger) (*Redis, error) {
	host := config.GetString("redis.url")
	password := config.GetString("redis.password")
	logger.Infof("Redis URI: %s", host)
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return &Redis{client}, nil
}
