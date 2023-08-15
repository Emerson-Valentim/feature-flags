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
	Key   string
	Value []byte
}

func (conn *RedisConnection) Set(input SetInput) error {
	err := conn.rdb.Set(conn.ctx, input.Key, input.Value, 0).Err()

	if err != nil {
		return errors.New("failed to set value")
	}

	return nil
}

type GetInput struct {
	Key string
}

func (conn *RedisConnection) Get(input GetInput) ([]byte, error) {
	value, err := conn.rdb.Get(conn.ctx, input.Key).Result()

	if err != nil {
		msg := err.Error()

		isNotFound := msg == "redis: nil"

		if isNotFound {
			msg = "not found"
		}

		return nil, errors.New(msg)
	}

	return []byte(value), nil
}

type DeleteInput struct {
	Key string
}

func (conn *RedisConnection) Delete(input DeleteInput) error {
	result, err := conn.rdb.Del(conn.ctx, input.Key).Result()

	if err != nil {
		return errors.New("failed to delete value")
	}

	if result == 0 {
		return errors.New("not found")
	}

	return nil
}
