package grpc

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc/grpclog"

	"simpleagent/pkg/util/log"
)

const (
	timestampOffset = 22
)

type redirectLogger struct {
}

func newRedirectLogger() redirectLogger {
	return redirectLogger{}
}

func (l redirectLogger) Write(b []byte) (int, error) {
	msg := string(b)

	levelSepIndex := strings.Index(msg, ":")
	msg = msg[levelSepIndex+timestampOffset:]

	switch b[0] {
	case 'I':
		log.Info(msg)
	case 'W':
		log.Warn(msg)
	case 'E':
		log.Error(msg)
	case 'F':
		log.Fatal(msg)
	default:
		log.Info(msg)
	}

	return 0, nil
}

func NewLogger() grpclog.LoggerV2 {
	errorW := ioutil.Discard
	warningW := ioutil.Discard
	infoW := ioutil.Discard

	logLevel := strings.ToLower(os.Getenv("GRPC_GO_LOG_SEVERITY_LEVEL"))
	switch logLevel {
	case "", "error": // If env is unset, set level to ERROR.
		errorW = newRedirectLogger()
	case "warning":
		warningW = newRedirectLogger()
	case "info":
		infoW = newRedirectLogger()
	}

	var v int
	vLevel := os.Getenv("GRPC_GO_LOG_VERBOSITY_LEVEL")
	if vl, err := strconv.Atoi(vLevel); err == nil {
		v = vl
	}

	return grpclog.NewLoggerV2WithVerbosity(infoW, warningW, errorW, v)
}
