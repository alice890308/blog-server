package rediskit

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Addr     string `long:"addr" env:"ADDR" description:"the address of Redis" required:"true"`
	Password string `long:"password" env:"PASSWORD" description:"the password of Redis"`
	Database string `long:"database" env:"DATABASE" description:"the database of Redis"`
}

type RedisClient struct {
	*redis.Client
	closeFunc func()
}

func (c *RedisClient) Close() error {
	if c.closeFunc != nil {
		c.closeFunc()
	}

	return c.Client.Close()
}

func NewRedisClient(ctx context.Context, conf *RedisConfig) *RedisClient {

}
