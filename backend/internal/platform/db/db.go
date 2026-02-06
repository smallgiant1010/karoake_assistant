package db

import (
	"context"
	"fmt"
	"karoake_assistant/backend/internal/data/sqlc"
	"github.com/jackc/pgx/v5"
)

func NewDatabaseConnection(connectionString string) (*pgx.Conn, *sqlc.Queries, error) {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, nil, fmt.Errorf("error occured connecting to database: %v", err)
	}

	queries := sqlc.New(conn)
	return conn, queries, nil
}
