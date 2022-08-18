package db

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"

	ent "entgo.io/bug/ent/generated"
	_ "entgo.io/bug/ent/generated/runtime"
)

func NewDriver(url string) (*sql.DB, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewClient(url string) (*ent.Client, error) {
	driver, err := NewDriver(url)
	if err != nil {
		return nil, err
	}

	entDriver := entsql.OpenDB(dialect.Postgres, driver)
	return ent.NewClient(
		ent.Driver(entDriver),
	), nil
}

func NewClientFromDriver(driver *sql.DB) (*ent.Client, error) {
	entDriver := entsql.OpenDB(dialect.Postgres, driver)
	return ent.NewClient(
		ent.Driver(entDriver),
	), nil
}
