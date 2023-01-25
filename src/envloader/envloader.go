package envloader

import (
	"errors"
	"fmt"
	"os"
)

type Environment struct {
	rootDir   string
	staticDir string
	apiKey    string
}

func Load() *Environment {
	env := new(Environment)
	env.rootDir = envReader("STATIC_HOSTER_HOME", fmt.Sprintf("%s/static-hoster", os.Getenv("HOME")))
	if _, err := os.Stat(fmt.Sprintf("%s/config.yml", env.rootDir)); errors.Is(err, os.ErrNotExist) {
		env.staticDir = envReader("STATIC_HOSTER_HOSTED_DIR", fmt.Sprintf("%s/static-hoster/hosted", os.Getenv("HOME")))
		env.apiKey = os.Getenv("STATIC_HOSTER_API_KEY")
	} else if err == nil {
		//TODO: Read from config.yml
	}
	return env
}

func envReader(envVar string, defaultVal string) string {
	if os.Getenv(envVar) == "" {
		return defaultVal
	} else {
		return os.Getenv(envVar)
	}
}
