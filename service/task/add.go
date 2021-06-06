package task

import "fmt"

// Add task
func Add(task *TaskConfig) error {
	fmt.Printf("add task %s with weight %v\n", task.Name, task.Weight)
	fmt.Printf("tags: %v\n", task.Tags)
	return nil
}
