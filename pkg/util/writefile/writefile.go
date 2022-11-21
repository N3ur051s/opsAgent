package writefile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"opsAgent/pkg/util/filesystem"
	"opsAgent/pkg/util/log"
)

type WriteFile struct {
	Name    string
	Path    string
	Content string
}

var mutex sync.RWMutex

func (writeFile *WriteFile) Execute() (string, error) {
	if err := mkdir(writeFile.Path); err != nil {
		log.Errorf("Failed to mkdir [%s]: %s", writeFile.Path, err)
		return "", err
	}
	file := filepath.Join(writeFile.Path, writeFile.Name)
	if err := ioutil.WriteFile(file, []byte(writeFile.Content), 0o644); err != nil {
		return "", err
	}

	perms, err := filesystem.NewPermission()
	if err != nil {
		return "", err
	}

	if err := perms.RestrictAccessToUser(file); err != nil {
		log.Errorf("Failed to write [%s] acl: %s", file, err)
		return "", err
	}

	log.Infof("Wrote File Success: %s", file)
	return fmt.Sprintf("Wrote File Success: %s", file), nil
}

func mkdir(path string) error {
	if path == "" {
		return nil
	}

	mutex.Lock()
	defer mutex.Unlock()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0o755); err != nil {
			return err
		}
	}
	return nil
}
