package task

type Config struct {
	// for factory
	MaxRetryTimes      int   // maximum number of attempts
	BaseRetryBackOffMs int64 // Backoff time for first retry
	MaxRetryBackOffMs  int64 // Maximum backoff time for retries, default is 50 seconds
	// for worker
	MaxIoWorkerNum int // Maximum number of workers (number of goroutines)
	MaxTaskNum     int // Maximum number of tasks
	MaxBlockSec    int // maximum blocking time
}

func GetDefaultConfig() *Config {
	return &Config{
		MaxRetryTimes:      10,
		BaseRetryBackOffMs: 100,
		MaxRetryBackOffMs:  50 * 1000,
		MaxIoWorkerNum:     50,
		MaxTaskNum:         1000,
		MaxBlockSec:        60,
	}
}
