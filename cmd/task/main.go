package task

import "github.com/urfave/cli/v2"

// NewTaskCommand factory
func NewTaskCommand() *cli.Command {
	return &cli.Command{
		Name:  "task",
		Usage: "task manipulation",
		Subcommands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Set task unique name",
						EnvVars:  []string{"TASK_NAME"},
						Required: true,
					},
					&cli.Int64Flag{
						Name:        "weight",
						Aliases:     []string{"w"},
						Usage:       "Set task weight",
						EnvVars:     []string{"TASK_WEIGHT"},
						Value:       1,
						DefaultText: "set to 1",
					},
					&cli.StringSliceFlag{
						Name:    "tags",
						Usage:   "Set task tags",
						EnvVars: []string{"TASK_TAGS"},
					},
				},
				Action: add,
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Specify the task to complete",
						EnvVars:  []string{"TASK_NAME"},
						Required: true,
					},
				},
				Action: complete,
			},
		},
	}
}
