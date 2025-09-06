package builder

import (
	"context"
	"fmt"
	"log"

	"github.com/MoonMoon1919/mayi"
	"github.com/MoonMoon1919/mayi/pkg/results"
	"github.com/MoonMoon1919/mayi/pkg/rules"
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

func New(svc service.Service) *cli.Command {
	patternFlagName := "pattern"
	patternFlag := cli.StringFlag{
		Name:     patternFlagName,
		Usage:    "the pattern you would like to ignore",
		Required: true,
	}

	ownersSliceFlagName := "owner"
	ownersFlag := cli.StringSliceFlag{
		Name:     ownersSliceFlagName,
		Usage:    "owner to assign to the rule - may be used more than once",
		Required: true,
	}

	ownerFlagName := "owner"
	ownerFlag := cli.StringFlag{
		Name:     ownerFlagName,
		Usage:    "github alias or email",
		Required: true,
	}

	sourcePatternName := "source-pattern"
	sourcePatternFlag := cli.StringFlag{
		Name:     sourcePatternName,
		Usage:    "The pattern of the rule you're moving",
		Required: true,
	}

	destinationPatternName := "destination-pattern"
	destinationPatternFlag := cli.StringFlag{
		Name:     destinationPatternName,
		Usage:    "The pattern of the rule you'd like to move a rule before or after",
		Required: true,
	}

	directionName := "direction"
	directionFlag := cli.StringFlag{
		Name:     directionName,
		Usage:    "The direction to move the rule - 'before' or 'after'",
		Required: true,
	}

	fixName := "fix"
	fixFlag := cli.BoolFlag{
		Name:  fixName,
		Usage: "Set to automatically fix found conflicts and optimize the file",
	}

	maxName := "max"
	maxFixes := cli.IntFlag{
		Name:  maxName,
		Value: 20,
		Usage: "The number of attempted fixes the autofixer will perform before exiting",
	}

	actionName := "action"
	actionFlag := cli.StringFlag{
		Name:  actionName,
		Value: "include",
		Usage: "if you would like to ignore or allow the path (useful for exclusion) - either 'include' or 'exclude'",
	}

	cmd := &cli.Command{
		Name:  "mayi-cli",
		Usage: "Manage and search your codeowners files with ease",
		Commands: []*cli.Command{
			makePathCommand("create", "Create a new CODEOWNERS file", []cli.Flag{},
				func(path string, c *cli.Command) error {
					err := svc.Init(path)
					if err != nil {
						return fmt.Errorf("❌ Failed to create file: %w", err)
					}

					return nil
				},
			),
			{
				Name:  "add",
				Usage: "Add new rules and rule owners",
				Commands: []*cli.Command{
					makePathCommand("rule", "Add a new rule", []cli.Flag{&patternFlag, &ownersFlag, &actionFlag},
						func(path string, c *cli.Command) error {
							pattern := c.String(patternFlagName)
							owners := c.StringSlice(ownersSliceFlagName)
							action := c.String(actionName)

							parsedAction, err := rules.ActionFromString(action)
							if err != nil {
								return err
							}

							results, err := svc.AddRule(path, pattern, owners, parsedAction)
							if err != nil {
								return fmt.Errorf("❌ Failed to add rule: %w", err)
							}

							logActionResults(results)
							return nil
						},
					),
					makePathCommand("owner", "Add a new owner to a rule", []cli.Flag{&patternFlag, &ownerFlag},
						func(path string, c *cli.Command) error {
							pattern := c.String(patternFlagName)
							owner := c.String(ownerFlagName)

							results, err := svc.AddRuleOwner(path, pattern, owner)
							if err != nil {
								return fmt.Errorf("❌ Failed to add rule owner: %w", err)
							}

							logActionResults(results)
							return nil
						},
					),
				},
			},
			{
				Name:  "delete",
				Usage: "Delete rules and rule owners",
				Commands: []*cli.Command{
					makePathCommand("rule", "Delete an existing rule", []cli.Flag{&patternFlag, &actionFlag},
						func(path string, c *cli.Command) error {
							pattern := c.String(patternFlagName)

							results, err := svc.RemoveRule(path, pattern)
							if err != nil {
								return fmt.Errorf("❌ Failed to remove rule: %w", err)
							}

							logActionResults(results)
							return nil
						},
					),
					makePathCommand("owner", "Remove an owner from a rule", []cli.Flag{&patternFlag, &ownerFlag},
						func(path string, c *cli.Command) error {
							pattern := c.String(patternFlagName)
							owner := c.String(ownerFlagName)

							results, err := svc.RemoveRuleOwner(path, pattern, owner)
							if err != nil {
								return fmt.Errorf("❌ Failed to remove rule owner: %w", err)
							}

							logActionResults(results)
							return nil
						},
					),
				},
			},
			{
				Name:  "get",
				Usage: "Search existing rules",
				Commands: []*cli.Command{
					makePathCommand("owners", "List owners for a rule", []cli.Flag{&patternFlag},
						func(path string, c *cli.Command) error {
							pattern := c.String(patternFlagName)

							owners, err := svc.GetOwnersForPath(path, pattern)
							if err != nil {
								return fmt.Errorf("❌ Failed to get owners for pattern: %w", err)
							}

							for _, owner := range owners {
								fmt.Printf("%s\n", owner)
							}

							return nil
						},
					),
				},
			},
			makePathCommand("move", "Move two existing rules", []cli.Flag{&sourcePatternFlag, &destinationPatternFlag, &directionFlag},
				func(path string, c *cli.Command) error {
					sourcePattern := c.String(sourcePatternName)
					destinationPattern := c.String(destinationPatternName)
					direction := c.String(directionName)

					parsedDirection, err := mayi.MoveDirectionFromString(direction)
					if err != nil {
						return fmt.Errorf("❌ Failed to parse direction: %w", err)
					}

					results, err := svc.MoveRule(path, sourcePattern, destinationPattern, parsedDirection)
					if err != nil {
						return fmt.Errorf("❌ Failed to move rule: %w", err)
					}

					logActionResults(results)
					return nil
				},
			),
			makePathCommand("analyze", "Check if your codeowners file has any conflicts, optionally fix them", []cli.Flag{&fixFlag, &maxFixes},
				func(path string, c *cli.Command) error {
					fix := c.Bool(fixName)
					maxFixes := c.Int(maxName)

					conflicts, err := svc.AnalyzeConflicts(path)
					if err != nil {
						return fmt.Errorf("❌ Failed to analyze conflicts %w", err)
					}

					if len(conflicts) == 0 {
						log.Print("No conflicts found")
						return nil
					}

					for _, conflict := range conflicts {
						log.Printf("FOUND CONFLICT: Left: %s, Right: %s, Type: %s \n", conflict.Left.Render(), conflict.Right.Render(), conflict.ConflictType)
					}

					// No need to check of we have more than 0 conflicts
					// since we're returning early in that case
					if fix {
						results, err := svc.FixConflicts(path, maxFixes)
						if err != nil {
							return fmt.Errorf("❌ Failed to fix conflicts %w", err)
						}
						logActionResults(results)
					}

					return nil
				},
			),
		},
	}

	return cmd
}
