package internal

import (
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"

	"simpleagent/pkg/version"
)

type ReadWaitCloser struct {
	pipeReader *io.PipeReader
	wg         sync.WaitGroup
}

func ProductToken() string {
	return fmt.Sprintf("simpleagent/%s Go/%s",
		version.AgentVersion, strings.TrimPrefix(runtime.Version(), "go"))
}
