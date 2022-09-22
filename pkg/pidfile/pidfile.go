package pidfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func WritePID(pidFilePath string) error {
	if byteContent, err := ioutil.ReadFile(pidFilePath); err == nil {
		pidStr := strings.TrimSpace(string(byteContent))
		pid, err := strconv.Atoi(pidStr)
		if err == nil && isProcess(pid) {
			return fmt.Errorf("Pidfile already exists, please check %s isn't running or remove %s",
				os.Args[0], pidFilePath)
		}
	}

	if err := os.MkdirAll(filepath.Dir(pidFilePath), os.FileMode(0755)); err != nil {
		return err
	}

	pidStr := fmt.Sprintf("%d", os.Getpid())
	if err := ioutil.WriteFile(pidFilePath, []byte(pidStr), 0644); err != nil {
		return err
	}

	return nil
}
