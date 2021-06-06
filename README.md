# Go CLI Example
A Go CLI example that shows a simple way to create CLI program with clean project structure.
## Installation
You could build from source:
```bash
git clone https://github.com/minghsu0107/go-cli-example.git
cd go-cli-example
go build -o gce
```
## Getting Started
We have two services, `task` and `template`. Each service is a subcommand with its own set of actions and flags. Take `task` service as an example. There are two actions for `task` service, `add` and `complete`, in which `add` adds a new task and `complete` completes the task. To add a task named `mytask` with a weight of 3, for example, you could run:
```bash
./gce task add --name=mytask --weight 3
```
To add additional tags, you could run:
```bash
./gce task add --name=mytask --weight=3 --tags=go,cli,example
```
You could run `./gce --help` to see help and full options.

Below is an overall project structure to see how commands are organized:
```
go-cli-example/
├── cmd
│   ├── task
│   │   ├── add.go
│   │   ├── complete.go
│   │   └── main.go
│   └── template
│       ├── add.go
│       ├── main.go
│       └── remove.go
├── go.mod
├── go.sum
├── main.go
└── service
    ├── task
    │   ├── add.go
    │   ├── complete.go
    │   └── config.go
    └── template
        ├── add.go
        ├── config.go
        └── remove.go
```