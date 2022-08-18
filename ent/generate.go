package ent

import (
	_ "ariga.io/atlas/sql/sqltool"
	_ "entgo.io/contrib/entgql"
	_ "entgo.io/ent"
	_ "github.com/golang-migrate/migrate/v4"
)

//go:generate go run entc.go
