package task

import "fmt"

// Complete task
func Complete(task *TaskConfig) error {
	fmt.Printf("completing task %s\n", task.Name)
	task.Complete = true
	fmt.Printf("is completed: %v\n", task.Complete)
	return nil
}
