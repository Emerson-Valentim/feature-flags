package repository

import (
	"encoding/json"
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
	return func(flag FlagEntity) (*FlagEntity, error) {
		encodedFlag, err := json.Marshal(flag)

		if err != nil {
			return nil, err
		}

		_, err = rdb.Get(redis.GetInput{
			Key: flag.Id,
		})

		if err == nil {
			return nil, err
		}

		rdb.Set(redis.SetInput{
			Key:   flag.Id,
			Value: encodedFlag,
		})

		return &flag, nil
	}
}

func find(rdb *redis.RedisConnection) FindFun {
	return func(ids []string) ([]FlagEntity, error) {
		var flags []FlagEntity

		for _, value := range ids {

			cachedFlag, err := rdb.Get(redis.GetInput{
				Key: value,
			})

			var decodedFlag FlagEntity

			json.Unmarshal(cachedFlag, &decodedFlag)

			if err != nil {
				return nil, err
			}

			flags = append(flags, decodedFlag)
		}

		return flags, nil
	}
}

func delete(rdb *redis.RedisConnection) DeleteFun {
	return func(id string) error {
		err := rdb.Delete(redis.DeleteInput{
			Key: id,
		})

		if err != nil {
			return err
		}

		return nil
	}
}

func update(rdb *redis.RedisConnection) UpdateFun {
	return func(flag FlagEntity) (*FlagEntity, error) {
		cachedFlag, err := rdb.Get(redis.GetInput{
			Key: flag.Id,
		})

		if err != nil {
			return nil, err
		}

		var decodedFlag FlagEntity

		json.Unmarshal(cachedFlag, &decodedFlag)

		updatedName := flag.Name
		updatedState := flag.State

		log.Println("here", updatedName, len(updatedName) == 0)

		if len(updatedName) == 0 {
			updatedName = decodedFlag.Name
		}

		if updatedState == nil {
			updatedState = decodedFlag.State
		}

		println(updatedState)

		updatedFlag := FlagEntity{
			Id:    decodedFlag.Id,
			Name:  updatedName,
			State: updatedState,
		}

		encodedFlag, err := json.Marshal(updatedFlag)

		if err != nil {
			return nil, err
		}

		rdb.Set(redis.SetInput{
			Key:   flag.Id,
			Value: encodedFlag,
		})

		return &updatedFlag, nil
	}
}
