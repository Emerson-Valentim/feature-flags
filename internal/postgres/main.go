package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Client struct {
	Conn *pgx.Conn
}

var (
	ErrRowNotFound = pgx.ErrNoRows
)

func NewPostgres(ctx context.Context, databaseURL string) *Client {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return nil
	}
	// defer conn.Close(context.Background())

	return &Client{
		Conn: conn,
	}
}
