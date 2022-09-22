package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"simpleagent/cmd/simpleagent/api"
	"simpleagent/cmd/simpleagent/common"
	"simpleagent/cmd/simpleagent/common/signals"
	. "simpleagent/conf"
	"simpleagent/pkg/pidfile"
	fyruntime "simpleagent/pkg/runtime"
	"simpleagent/pkg/util/log"
)

var (
	pidfilePath string

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the Agent",
		Long:  `Runs the agent in the foreground`,
		RunE:  run,
	}
)

func init() {
	AgentCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&pidfilePath, "pidfile", "p", "", "path to the pidfile")
}

func run(cmd *cobra.Command, args []string) error {
	defer func() {
		StopAgent()
	}()

	fyruntime.SetMaxProcs()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	stopCh := make(chan error)

	go func() {
		select {
		case <-signals.Stopper:
			log.Info("Received stop command, shutting down...")
			stopCh <- nil
		case <-signals.ErrorStopper:
			log.Fatal("The Agent has encountered an error, shutting down...")
			stopCh <- fmt.Errorf("shutting down because of an error")
		case sig := <-signalCh:
			log.Infof("Received signal '%s', shutting down...", sig)
			stopCh <- nil
		}
	}()

	// intercept the SIGPIPE signals and discard them.
	sigpipeCh := make(chan os.Signal, 1)
	signal.Notify(sigpipeCh, syscall.SIGPIPE)
	go func() {
		for range sigpipeCh {
			// do nothing
		}
	}()

	if err := StartAgent(); err != nil {
		return err
	}

	select {
	case err := <-stopCh:
		return err
	}
}

func StartAgent() error {
	var (
		err            error
		configSetupErr error
	)

	common.MainCtx, common.MainCtxCancel = context.WithCancel(context.Background())

	configSetupErr = LoadConfig(confFilePath)
	if configSetupErr != nil {
		log.Errorf("Failed to setup config %v", configSetupErr)
		return fmt.Errorf("unable to set up global agent configuration: %v", configSetupErr)
	}

	if pidfilePath != "" {
		err = pidfile.WritePID(pidfilePath)
		if err != nil {
			return fmt.Errorf("Error while writing PID file, exiting: %v", err)
		}
		log.Infof("pid '%d' written to pid file '%s'", os.Getpid(), pidfilePath)
	}

	if err = api.StartServer(*Conf.Server); err != nil {
		return fmt.Errorf("Error while starting agent, exiting: %v", err)
	}

	return nil
}

func StopAgent() {
	api.StopServer()

	os.Remove(pidfilePath)

	// gracefully shut down any component
	common.MainCtxCancel()

	log.Info("See ya!")
}
