package main

import (
	"fmt"
	"log"
	"os"

	"github.com/minghsu0107/go-cli-example/cmd/task"
	"github.com/minghsu0107/go-cli-example/cmd/template"
	"github.com/urfave/cli/v2"
)

var (
	version = "0.0.0"
	build   = "unkown"
)

func main() {
	app := cli.NewApp()
	app.Name = "Task CLI"
	app.Usage = "Task CLI"
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Commands = append(app.Commands,
		task.NewTaskCommand(),
		template.NewTemplateCommand(),
	)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
