//go:build ignore

package main

import (
	"path/filepath"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/sirupsen/logrus"
)

const schemaPath = "./schema"
const generatedPath = "./generated"

func main() {
	entgqlExt, err := entgql.NewExtension()
	if err != nil {
		logrus.WithError(err).Fatalln("error creating entgql extension")
	}

	schemaPathAbsDir, err := filepath.Abs(schemaPath)
	if err != nil {
		logrus.WithError(err).Fatalln("error getting absolute path of `schemaPath`")
	}

	generatedPathAbsDir, err := filepath.Abs(generatedPath)
	if err != nil {
		logrus.WithError(err).Fatalln("error getting absolute path of `generatedPath`")
	}

	schemaGraph, err := entc.LoadGraph(schemaPathAbsDir, &gen.Config{})
	if err != nil {
		logrus.WithError(err).Fatalln("error loading schema graph")
	}

	if err := entc.Generate(
		schemaPathAbsDir,
		&gen.Config{
			Target:  generatedPathAbsDir,
			Package: schemaGraph.Config.Package + "/" + filepath.Clean(generatedPath),
		},
		entc.FeatureNames("entql"),
		entc.FeatureNames("privacy"),
		entc.FeatureNames("schema/snapshot"),
		entc.FeatureNames("sql/schemaconfig"),
		entc.FeatureNames("sql/upsert"),
		entc.FeatureNames("sql/versioned-migration"),
		entc.Extensions(entgqlExt),
	); err != nil {
		logrus.WithError(err).Fatalln("error running ent codegen")
	}
}
