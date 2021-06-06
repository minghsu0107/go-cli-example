package task

import (
	"github.com/minghsu0107/go-cli-example/service/task"
	"github.com/urfave/cli/v2"
)

func add(c *cli.Context) error {
	config := &task.TaskConfig{
		Name:   c.String("name"),
		Weight: c.Int64("weight"),
		Tags:   c.StringSlice("tags"),
	}
	return task.Add(config)
}
