package conf

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

var (
	Conf config
)

type config struct {
	LogLevel string `toml:"loglevel"`
	Node     string `toml:"node"`
	Server   *Server
	Opsfast  *Opsfast `toml:"ops-fast"`
}

type Server struct {
	Hostname string `toml:"hostname"`
	Addr     string `toml:"addr"`
	Port     int64  `toml:"port"`
}

type Opsfast struct {
	Url      string `toml:"url"`
	ApiToken string `toml:"apiToken"`
}

func getDefaultConfigPath() (string, error) {
	envfile := os.Getenv("Agent_CONFIG_PATH")
	homefile := os.ExpandEnv("${HOME}/.opsAgent/.conf")
	etcfile := "/etc/opsAgent/opsAgent.conf"
	for _, path := range []string{envfile, homefile, etcfile} {
		if _, err := os.Stat(path); err == nil {
			log.Infof("Using config file: %s", path)
			return path, nil
		}
	}

	return "", fmt.Errorf("No config file specified, and could not find one"+
		" in $opsAgent_CONFIG_PATH, %s, or %s", homefile, etcfile)
}

func LoadConfig(path string) error {
	var err error
	if path == "" {
		if path, err = getDefaultConfigPath(); err != nil {
			return err
		}
	}
	data, err := LoadConfigFile(path)
	if err != nil {
		return fmt.Errorf("Error loading config file %s: %w", path, err)
	}

	_, err = toml.Decode(string(data), &Conf)
	if err != nil {
		return fmt.Errorf("config decode err:%v", err)
	}

	log.Infof("config data:%v", Conf)
	return nil
}

func LoadConfigFile(config string) ([]byte, error) {
	buffer, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func GetLogLvl() log.Level {
	// DEBUG INFO WARN ERROR OFF
	switch Conf.LogLevel {
	case "DEBUG":
		return log.DebugLevel
	case "INFO":
		return log.InfoLevel
	case "WARN":
		return log.WarnLevel
	case "ERROR":
		return log.ErrorLevel
	case "FATAL":
		return log.FatalLevel
	}

	return log.DebugLevel
}
