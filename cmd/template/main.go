package template

import "github.com/urfave/cli/v2"

// NewTemplatCommands factory
func NewTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:  "template",
		Usage: "template manipulation",
		Subcommands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a template",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Set template unique name",
						EnvVars:  []string{"TEMPLATE_NAME"},
						Required: true,
					},
					&cli.Int64Flag{
						Name:        "weight",
						Aliases:     []string{"w"},
						Usage:       "Set template weight",
						EnvVars:     []string{"TEMPLATE_WEIGHT"},
						Value:       1,
						DefaultText: "set to 1",
					},
					&cli.StringSliceFlag{
						Name:    "tags",
						Usage:   "Set template tags",
						EnvVars: []string{"TEMPLATE_TAGS"},
					},
				},
				Action: add,
			},
			{
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "remove a template",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Specify the template to remove",
						EnvVars:  []string{"TEMPLATE_NAME"},
						Required: true,
					},
				},
				Action: remove,
			},
		},
	}
}
