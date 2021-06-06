package template

import (
	"github.com/minghsu0107/go-cli-example/service/template"

	"github.com/urfave/cli/v2"
)

func remove(c *cli.Context) error {
	config := &template.TemplateConfig{
		Name:   c.String("name"),
		Weight: c.Int64("weight"),
		Tags:   c.StringSlice("tags"),
	}
	return template.Remove(config)
}
