package main

import (
	"context"
	"log"
	"os"

	"github.com/MoonMoon1919/mayi"
	"github.com/MoonMoon1919/mayi-cli/internal/builder"
	"github.com/MoonMoon1919/mayi/pkg/service"
)

func run(args []string, svc service.Service) {
	cmd := builder.New(svc)

	if err := cmd.Run(context.Background(), args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	repo := service.NewFileRepository(
		mayi.RenderOptions{
			TrailingNewLine: true,
			HeaderComment:   "This file was automatically generated, do not make manual edits",
		},
	)

	svc := service.New(repo)

	run(os.Args, svc)
}
