package repository

import (
	"feature-flag/go/ports/redis"
	"log"
)

func NewRedisAdapter(host string) (Repository, error) {
	rdb, err := redis.NewRedisConnection(host)

	if err != nil {
		log.Fatal("Failed to create redis connection")
	}

	return Repository{
		Insert: insert(rdb),
		Find:   find(rdb),
		Delete: delete(rdb),
		Update: update(rdb),
	}, nil
}

func insert(rdb *redis.RedisConnection) InsertFun {
	return func(flag FlagEntity) (FlagEntity, error) {
		log.Printf("Inserting")
		return flag, nil
	}
}

func find(rdb *redis.RedisConnection) FindFun {
	return func(ids []string) ([]FlagEntity, error) {
		log.Printf("Finding")
		return nil, nil
	}
}

func delete(rdb *redis.RedisConnection) DeleteFun {
	return func(id string) error {
		log.Printf("Deleting")
		return nil
	}
}

func update(rdb *redis.RedisConnection) UpdateFun {
	return func(flag FlagEntity) (FlagEntity, error) {
		log.Printf("Updating")
		return flag, nil
	}
}
