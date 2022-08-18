//go:build ignore

package main

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/sirupsen/logrus"

	"entgo.io/bug/ent/db"
	"entgo.io/bug/ent/generated/migrate"
)

const migrationsPath = "./migrations"

func main() {
	databaseURL := "postgres://postgres:pass@localhost:5434/test"

	migrationsPathAbsDir, err := filepath.Abs(migrationsPath)
	if err != nil {
		logrus.WithError(err).Fatalln("error getting absolute path of `migrationsPath`")
	}
	if _, err := os.Stat(migrationsPathAbsDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(migrationsPathAbsDir, os.ModePerm)
		if err != nil {
			logrus.WithError(err).Fatalln("error creating migrations directory")
		}
	}

	driver, err := db.NewDriver(databaseURL)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer driver.Close()

	dir, err := sqltool.NewGolangMigrateDir(migrationsPathAbsDir)
	if err != nil {
		logrus.Fatalln(err)
	}

	ctx := context.Background()
	if err := migrate.NewSchema(
		entsql.OpenDB(dialect.Postgres, driver),
	).Diff(
		ctx,
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
	); err != nil {
		logrus.Fatalf("failed creating schema resources: %v", err)
	}
}
