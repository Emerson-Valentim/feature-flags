package redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisConnection(host string) (*RedisConnection, error) {
	opts, err := redis.ParseURL(host)

	if err != nil {
		return nil, errors.New("Invalid connection URL")
	}

	ctx := context.Background()
	rdb := redis.NewClient(opts)

	return &RedisConnection{rdb, ctx}, nil
}

type SetInput struct {
	key   string
	value []byte
}

func (conn *RedisConnection) Set(input SetInput) error {
	err := conn.rdb.Set(conn.ctx, input.key, input.value, 0).Err()

	if err != nil {
		return errors.New("Failed to set value")
	}

	return nil
}

type GetInput struct {
	key string
}

func (conn *RedisConnection) Get(input GetInput) ([]byte, error) {
	value, err := conn.rdb.Get(conn.ctx, input.key).Result()

	if err != nil {
		return nil, errors.New("Failed to get value")
	}

	return []byte(value), nil
}

type DeleteInput struct {
	key string
}

func (conn *RedisConnection) Delete(input DeleteInput) error {
	err := conn.rdb.FunctionDelete(conn.ctx, input.key).Err()

	if err != nil {
		return errors.New("Failed to delete value")
	}

	return nil
}
