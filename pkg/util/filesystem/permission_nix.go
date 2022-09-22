package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strconv"

	"simpleagent/pkg/util/log"
)

type Permission struct{}

func NewPermission() (*Permission, error) {
	return &Permission{}, nil
}

func (p *Permission) RestrictAccessToUser(path string) error {
	usr, err := user.Lookup("fy-agent")
	if err != nil {
		return nil
	}

	usrID, err := strconv.Atoi(usr.Uid)
	if err != nil {
		return fmt.Errorf("couldn't parse UID (%s): %w", usr.Uid, err)
	}

	grpID, err := strconv.Atoi(usr.Gid)
	if err != nil {
		return fmt.Errorf("couldn't parse GID (%s): %w", usr.Gid, err)
	}

	if err = os.Chown(path, usrID, grpID); err != nil {
		if errors.Is(err, fs.ErrPermission) {
			log.Infof("Cannot change owner of '%s', permission denied", path)
			return nil
		}

		return fmt.Errorf("couldn't set user and group owner for %s: %w", path, err)
	}

	return nil
}
