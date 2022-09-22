package task

type Task interface {
	Execute() (string, error)
	CallBack(out string, result *Result, err error)
}

// task is used in the project
type task struct {
	attemptCount         int     // number of attempts
	maxRetryTimes        int     // maximum number of attempts
	baseRetryBackOffMs   int64   // Backoff time for first retry
	maxRetryIntervalInMs int64   // Maximum backoff time for retries, default is 50 seconds
	createTimeMs         int64   // time of creation
	nextRetryMs          int64   // next retry time
	result               *Result // send result
	task                 Task
}

// *************************
// task factory
type taskFactory struct {
	maxRetryTimes        int   // maximum number of attempts
	baseRetryBackOffMs   int64 // Backoff time for first retry
	maxRetryIntervalInMs int64 // Maximum backoff time for retries, default is 50 seconds
}

func newTaskFactory(c *Config) *taskFactory {
	return &taskFactory{
		maxRetryTimes:        c.MaxRetryTimes,
		baseRetryBackOffMs:   c.BaseRetryBackOffMs,
		maxRetryIntervalInMs: c.MaxRetryBackOffMs,
	}
}

func (taskFactory *taskFactory) produce(t Task) *task {
	return &task{
		attemptCount:         0,
		maxRetryTimes:        taskFactory.maxRetryTimes,
		baseRetryBackOffMs:   taskFactory.baseRetryBackOffMs,
		maxRetryIntervalInMs: taskFactory.maxRetryIntervalInMs,
		createTimeMs:         getTimeMs(),
		nextRetryMs:          0,
		result:               initResult(t),
		task:                 t,
	}
}
