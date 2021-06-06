package task

// TaskConfig config
type TaskConfig struct {
	Name     string
	Weight   int64
	Complete bool
	Tags     []string
}
