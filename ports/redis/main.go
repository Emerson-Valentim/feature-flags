package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
	conn *redis.Client
	ctx  context.Context
}

func NewRedisConnection(host string) RedisConnection {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: host,
	})

	return RedisConnection{rdb, ctx}
}

func (conn *redis.Client) Set() {
	err := conn.Set()
}
