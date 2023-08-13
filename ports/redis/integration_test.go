package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var redisConnectionString = "redis://localhost:6379"

func TestNewRedisConnection_InvalidURL(t *testing.T) {
	conn, err := NewRedisConnection("invalid-url")
	assert.Nil(t, conn)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid connection URL", err.Error())
}

func TestNewRedisConnection_ValidURL(t *testing.T) {
	conn, err := NewRedisConnection(redisConnectionString)
	assert.NotNil(t, conn)
	assert.Nil(t, err)
}

func TestSet(t *testing.T) {
	conn, err := NewRedisConnection(redisConnectionString)
	assert.NotNil(t, conn)
	assert.Nil(t, err)
	defer conn.rdb.Close()

	key := "test-key"
	value := []byte("test-value")

	err = conn.Set(SetInput{key: key, value: value})
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	conn, err := NewRedisConnection(redisConnectionString)
	assert.NotNil(t, conn)
	assert.Nil(t, err)
	defer conn.rdb.Close()

	key := "test-key"
	value := []byte("test-value")

	retrievedValue, err := conn.Get(GetInput{key: key})
	assert.Nil(t, err)
	assert.Equal(t, value, retrievedValue)
}

func TestDelete(t *testing.T) {
	conn, err := NewRedisConnection(redisConnectionString)
	assert.NotNil(t, conn)
	assert.Nil(t, err)
	defer conn.rdb.Close()

	key := "test-key"

	err = conn.Delete(DeleteInput{key: key})
	assert.Nil(t, err)

	_, err = conn.Get(GetInput{key: key})
	assert.NotNil(t, err)
}
