package task

import (
	"github.com/minghsu0107/go-cli-example/service/task"
	"github.com/urfave/cli/v2"
)

func complete(c *cli.Context) error {
	config := &task.TaskConfig{
		Name: c.String("name"),
	}
	return task.Complete(config)
}
