package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RunDB() error {
	pgUrl := "postgres://docker:password@localhost:5432/appointments"
	db, err := pgxpool.New(context.Background(), pgUrl)

	if err != nil {
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("Database is connected")
	return nil
}
