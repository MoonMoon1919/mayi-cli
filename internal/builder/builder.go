package builder

import (
	"github.com/MoonMoon1919/mayi/pkg/service"
	"github.com/urfave/cli/v3"
)

func New(svc service.Service) *cli.Command {
	cmd := &cli.Command{
		Name:     "mayi-cli",
		Usage:    "Manage and search your codeowners files with ease",
		Commands: []*cli.Command{},
	}

	return cmd
}
