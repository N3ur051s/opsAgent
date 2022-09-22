package hostname

import (
	"context"
	"os/exec"
	"strings"
	"time"
)

func getSystemFQDN() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/hostname", "-f")

	out, err := cmd.Output()
	return strings.TrimSpace(string(out)), err
}
