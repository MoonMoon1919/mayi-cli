package main

import (
	"log"
	"os"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute/pkg/app"
	"github.com/MoonMoon1919/mayi-cli/docs/internal/documents"
)

func main() {
	repo := doyoucompute.NewFileRepository()
	fileRenderer := doyoucompute.NewMarkdownRenderer()
	execRenderer := doyoucompute.NewExecutionRenderer()
	runner := doyoucompute.NewTaskRunner(doyoucompute.DefaultSecureConfig())
	svc := doyoucompute.NewService(repo, runner, fileRenderer, execRenderer)

	app := app.New(&svc)

	// Load & register docs
	readme, err := documents.ReadMe()
	if err != nil {
		log.Fatal(err)
	}

	app.Register(readme)

	contrib, err := documents.Contributing()
	if err != nil {
		log.Fatal(err)
	}
	app.Register(contrib)

	bugreport, err := documents.BugReport()
	if err != nil {
		log.Fatal(err)
	}
	app.Register(bugreport)

	pullrequest, err := documents.PullRequest()
	if err != nil {
		log.Fatal(err)
	}
	app.Register(pullrequest)

	app.Run(os.Args)
}
