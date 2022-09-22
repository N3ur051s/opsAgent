package conf

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"

	"simpleagent/internal"
)

var (
	Conf config

	httpLoadConfigRetryInterval = 10 * time.Second

	// fetchURLRe is a regex to determine whether the requested file should
	// be fetched from a remote or read from the filesystem.
	fetchURLRe = regexp.MustCompile(`^\w+://`)
)

type config struct {
	LogLevel string `toml:"loglevel"`
	Server   *Server
}

type Server struct {
	Hostname string `toml:"hostname"`
	Addr     string `toml:"addr"`
	Port     int64  `toml:"port"`
}

func getDefaultConfigPath() (string, error) {
	envfile := os.Getenv("simpleagent_CONFIG_PATH")
	homefile := os.ExpandEnv("${HOME}/.simpleagent/.conf")
	etcfile := "/etc/simpleagent/simpleagent.conf"
	if runtime.GOOS == "windows" {
		programFiles := os.Getenv("ProgramFiles")
		if programFiles == "" { // Should never happen
			programFiles = `C:\Program Files`
		}
		etcfile = programFiles + `\simpleagent\simpleagent.conf`
	}
	for _, path := range []string{envfile, homefile, etcfile} {
		if isURL(path) {
			log.Printf("I! Using config url: %s", path)
			return path, nil
		}
		if _, err := os.Stat(path); err == nil {
			log.Printf("I! Using config file: %s", path)
			return path, nil
		}
	}

	return "", fmt.Errorf("No config file specified, and could not find one"+
		" in $simpleagent_CONFIG_PATH, %s, or %s", homefile, etcfile)
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
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
	if fetchURLRe.MatchString(config) {
		u, err := url.Parse(config)
		if err != nil {
			return nil, err
		}

		switch u.Scheme {
		case "https", "http":
			return fetchConfig(u)
		default:
			return nil, fmt.Errorf("scheme %q not supported", u.Scheme)
		}
	}

	// If it isn't a https scheme, try it as a file
	buffer, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}

	mimeType := http.DetectContentType(buffer)
	if !strings.Contains(mimeType, "text/plain") {
		return nil, fmt.Errorf("provided config is not a TOML file: %s", config)
	}

	return buffer, nil
}

func fetchConfig(u *url.URL) ([]byte, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	if v, exists := os.LookupEnv("simpleagent_TOKEN"); exists {
		req.Header.Add("Authorization", "Token "+v)
	}
	req.Header.Add("Accept", "application/toml")
	req.Header.Set("User-Agent", internal.ProductToken())

	retries := 3
	for i := 0; i <= retries; i++ {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("Retry %d of %d failed connecting to HTTP config server %s", i, retries, err)
		}

		if resp.StatusCode != http.StatusOK {
			if i < retries {
				log.Printf("Error getting HTTP config.  Retry %d of %d in %s.  Status=%d", i, retries, httpLoadConfigRetryInterval, resp.StatusCode)
				time.Sleep(httpLoadConfigRetryInterval)
				continue
			}
			return nil, fmt.Errorf("Retry %d of %d failed to retrieve remote config: %s", i, retries, resp.Status)
		}
		defer resp.Body.Close()
		return io.ReadAll(resp.Body)
	}

	return nil, nil
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
