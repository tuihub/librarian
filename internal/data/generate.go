//go:build ignore

package main

//go:generate go run generate.go -out ./internal/query

import (
	"flag"

	"github.com/tuihub/librarian/internal/data/internal/models"

	"gorm.io/gen"
)

func main() {
	outPath := flag.String("out", "./internal/query", "output path for generated query code")
	flag.Parse()

	// 1. Generate Query Code
	g := gen.NewGenerator(gen.Config{
		OutPath:       *outPath,
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.ApplyBasic(models.GetModels()...)
	g.Execute()
}
