package ent

import (
	"embed"
	"io/fs"

	"github.com/sirupsen/logrus"
)

var (
	//go:embed migrations/*
	Migrations embed.FS
)

func init() {
	chrootFS, err := fs.Sub(Migrations, "migrations")
	if err != nil {
		logrus.WithError(err).Fatal("error initializing sql migrations")
	}

	if assertFS, ok := chrootFS.(embed.FS); ok {
		Migrations = assertFS
	}
}
