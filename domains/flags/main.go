package flags

import (
	"feature-flag/go/domains/flags/repository"
	"log"

	"github.com/google/uuid"
)

type Flag struct {
	Id    string
	Name  string
	State bool
}

type CreateInput struct {
	Name string
}
type CreateFun func(input CreateInput) (Flag, error)

type GetFun func(id string) (Flag, error)

type UpdateInput struct {
	State bool
	Name  string
}
type UpdateFun func(id string, input UpdateInput) (Flag, error)

type DeleteFun func(id string) error

type FlagsDomain struct {
	Create CreateFun
	Get    GetFun
	Update UpdateFun
	Delete DeleteFun
}

func NewFlags(host string) (FlagsDomain, error) {
	R, err := repository.NewRedisAdapter(host)

	if err != nil {
		log.Fatal("Failed to create flags domain")
	}

	return Flags(&R), nil
}

func Flags(R *repository.Repository) FlagsDomain {
	return FlagsDomain{
		Create: create(R),
		Get:    get(R),
		Update: update(R),
		Delete: delete(R),
	}
}

func create(R *repository.Repository) CreateFun {
	return func(input CreateInput) (Flag, error) {
		id := uuid.New().String()
		name := input.Name
		state := false

		databaseFlag, err := R.Insert(repository.FlagEntity{
			Id:    id,
			Name:  name,
			State: state,
		})

		flag := Flag{
			Id:    databaseFlag.Id,
			Name:  databaseFlag.Name,
			State: databaseFlag.State,
		}

		if err != nil {
			return flag, err
		}

		return flag, nil
	}
}

func get(R *repository.Repository) GetFun {
	return func(id string) (Flag, error) {
		databaseFlags, err := R.Find([]string{id})

		flag := Flag{
			Id:    databaseFlags[0].Id,
			Name:  databaseFlags[0].Name,
			State: databaseFlags[0].State,
		}

		if err != nil {
			return flag, err
		}

		return flag, nil
	}
}

func update(R *repository.Repository) UpdateFun {
	return func(id string, input UpdateInput) (Flag, error) {
		databaseFlag, err := R.Update(repository.FlagEntity{
			Id:    id,
			Name:  input.Name,
			State: input.State,
		})

		flag := Flag{
			Id:    databaseFlag.Id,
			Name:  databaseFlag.Name,
			State: databaseFlag.State,
		}

		if err != nil {
			return flag, err
		}

		return flag, nil
	}
}

func delete(R *repository.Repository) DeleteFun {
	return func(id string) error {
		err := R.Delete(id)

		if err != nil {
			return err
		}

		return nil
	}
}
