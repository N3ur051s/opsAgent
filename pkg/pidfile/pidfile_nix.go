package pidfile

import (
	"os"
	"path/filepath"
	"strconv"
)

// isProcess searches for the PID under /proc
func isProcess(pid int) bool {
	_, err := os.Stat(filepath.Join("/proc", strconv.Itoa(pid)))
	return err == nil
}
