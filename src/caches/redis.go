package caches

import (
	"github.com/go-redis/redis"
	"github.com/gogotsenghsien/simple-rate-limit/src/configs"
)

type Redis struct {
	*redis.ClusterClient
}

func NewRedis(config *configs.Config) (*Redis, error) {
	host := config.GetString("redis.url")
	password := config.GetString("redis.password")
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{host},
		Password: password,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return &Redis{client}, nil
}
