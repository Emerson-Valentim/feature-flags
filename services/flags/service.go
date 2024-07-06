package flags

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (*Flag, error)
	Insert(ctx context.Context, flag Flag) error
}

type Service struct {
	repo Repository
}

var (
	ErrFlagNotFound = errors.New("flag not found")
)

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

type Flag struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Status bool      `json:"status"`
}

func (s *Service) Create(ctx context.Context, name string) (*Flag, error) {
	flag := Flag{
		ID:     uuid.New(),
		Name:   name,
		Status: true,
	}

	err := s.repo.Insert(ctx, flag)
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
		return nil, errors.New("failed to insert")
	}

	return &flag, nil
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*Flag, error) {
	flag, err := s.repo.Get(ctx, id)
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
		return nil, errors.Join(err)
	}

	return flag, nil
}
