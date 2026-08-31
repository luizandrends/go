package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	urlExample := "postgres://pg:password@localhost:5433/tests"
	db, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	query := "create table if not exists foo(id bigserial primary key, bar varchar(255));"
	if _, err := db.Exec(context.Background(), query); err != nil {
		panic(err)
	}

	query = "insert into foo (bar) values ($1)"
	if _, err := db.Exec(context.Background(), query, "abdef"); err != nil {
		panic(err)
	}

	query = "select * from foo limit 1;"

	type foobar struct {
		id  int64
		bar string
	}
	var res foobar

	if err := db.QueryRow(context.Background(), query).Scan(&res.id, &res.bar); err != nil {
		panic(err)
	}

	fmt.Printf("%#+v", res)
}
