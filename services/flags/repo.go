package flags

import (
	"context"
	"errors"
	"fmt"
	"main/internal/postgres"

	"github.com/google/uuid"
)

type PostgresRepository struct {
	pg *postgres.Client
}

func NewPostgresRepository(pg *postgres.Client) PostgresRepository {
	return PostgresRepository{
		pg: pg,
	}
}

func (p PostgresRepository) Insert(ctx context.Context, flag Flag) error {
	_, err := p.pg.Conn.Exec(ctx, `INSERT into "flags" ("id", "name", "status") values ($1, $2, $3)`, flag.ID.String(), flag.Name, flag.Status)
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
		return errors.New("failed to insert")
	}
	return nil
}

func (p PostgresRepository) Get(ctx context.Context, id uuid.UUID) (*Flag, error) {
	var flagID uuid.UUID
	var name string
	var status bool
	err := p.pg.Conn.QueryRow(ctx, `SELECT id, name, status from "flags" where id = $1`, id.String()).Scan(&flagID, &name, &status)
	switch {
	case errors.Is(postgres.ErrRowNotFound, err):
		return nil, ErrFlagNotFound
	case err != nil:
		fmt.Println()
		return nil, errors.New("failed to query")
	default:
		return &Flag{
			ID:     flagID,
			Name:   name,
			Status: status,
		}, nil
	}
}
