package runtime

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"go.uber.org/automaxprocs/maxprocs"

	"simpleagent/conf"
	"simpleagent/pkg/util/log"
)

const (
	gomaxprocsKey = "GOMAXPROCS"
)

func SetMaxProcs() {
	log.SetLevel(conf.GetLogLvl())
	defer func() {
		log.Infof("runtime: final GOMAXPROCS value is: %d", runtime.GOMAXPROCS(0))
	}()

	_, err := maxprocs.Set(maxprocs.Logger(log.Debugf))
	if err != nil {
		log.Errorf("runtime: error auto-setting maxprocs: %v ", err)
	}

	if max, exists := os.LookupEnv(gomaxprocsKey); exists {
		if max == "" {
			log.Errorf("runtime: GOMAXPROCS value was empty string")
			return
		}

		_, err = strconv.Atoi(max)
		if err == nil {
			return
		}

		if strings.HasSuffix(max, "m") {
			trimmed := strings.TrimSuffix(max, "m")
			milliCPUs, err := strconv.Atoi(trimmed)
			if err != nil {
				log.Errorf("runtime: error parsing GOMAXPROCS milliCPUs value: %v", max)
				return
			}

			cpus := milliCPUs / 1000
			if cpus > 0 {
				log.Infof("runtime: honoring GOMAXPROCS millicpu configuration: %v, setting GOMAXPROCS to: %d", max, cpus)
				runtime.GOMAXPROCS(cpus)
			} else {
				log.Infof(
					"runtime: GOMAXPROCS millicpu configuration: %s was less than 1, setting GOMAXPROCS to 1",
					max)
				runtime.GOMAXPROCS(1)
			}
			return
		}

		log.Errorf(
			"runtime: unhandled GOMAXPROCS value: %s", max)
	}
}

// NumVCPU returns the number of virtualizes CPUs available to the process. It should be used instead of
// runtime.NumCPU() in virtualized environments like K8s to ensure that processes don't attempt to
// over-subscribe CPUs. For example, on a 16 vCPU machine in a docker container allocated 8 vCPUs,
// runtime.NumCPU() will return 16 but NumVCPU() will return 8.
func NumVCPU() int {
	return runtime.GOMAXPROCS(0)
}
