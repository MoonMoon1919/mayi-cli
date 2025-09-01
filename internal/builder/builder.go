package builder

import (
	"context"
	"log"

	"github.com/MoonMoon1919/mayi/pkg/results"
	"github.com/MoonMoon1919/mayi/pkg/service"
	"github.com/urfave/cli/v3"
)

func logActionResults(results []results.Result) {
	for _, result := range results {
		log.Print(result.Log())
	}
}

func makePathCommand(name, usage string, flags []cli.Flag, handler func(path string, c *cli.Command) error) *cli.Command {
	pathFlag := cli.StringFlag{
		Name:  "path",
		Value: ".github/CODEOWNERS",
		Usage: "The path to which your CODEOWNERS file will be saved",
	}

	return &cli.Command{
		Name:  name,
		Usage: usage,
		Flags: append([]cli.Flag{&pathFlag}, flags...),
		Action: func(ctx context.Context, c *cli.Command) error {
			path := c.String("path")

			return handler(path, c)
		},
	}
}

func makeRuleCommand(name, usage string, flags []cli.Flag, handler func(path string, c *cli.Command) error) *cli.Command {
	return makePathCommand(name, usage, flags, func(path string, c *cli.Command) error {
		return handler(path, c)
	})
}

func New(svc service.Service) *cli.Command {
	filePathFlag := cli.StringFlag{
		Name:     "filepath",
		Usage:    "the filepath to file you would like to ignore",
		Required: true,
	}
	ownersFlag := cli.StringSliceFlag{
		Name:     "owners",
		Usage:    "list of owners to assign to the rule",
		Required: true,
	}

	cmd := &cli.Command{
		Name:  "mayi-cli",
		Usage: "Manage and search your codeowners files with ease",
		Commands: []*cli.Command{
			makePathCommand("create", "create a new CODEOWNERS file", []cli.Flag{},
				func(path string, c *cli.Command) error {
					return svc.Init(path)
				},
			),
			{
				Name:  "add",
				Usage: "Add a new rule",
				Commands: []*cli.Command{
					makeRuleCommand("file", "Add a new file rule", []cli.Flag{&filePathFlag, &ownersFlag},
						func(path string, c *cli.Command) error {
							filePath := c.String("filepath")
							owners := c.StringSlice("owners")

							results, err := svc.AddRule(path, filePath, owners)
							logActionResults(results)
							return err
						},
					),
				},
			},
			{
				Name:  "delete",
				Usage: "Delete an existing rule",
			},
			{
				Name:  "move",
				Usage: "Move two existing rules",
			},
			{
				Name:  "analyze",
				Usage: "Check if your codeowners file has any conflicts, optionally fix them",
			},
		},
	}

	return cmd
}
