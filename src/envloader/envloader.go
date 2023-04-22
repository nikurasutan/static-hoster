package envloader

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	RootDir   string
	StaticDir string
	User      string
	ApiKey    string
	Port      string
	BaseRoute string
}

func Load() (env *Environment) {
	err := godotenv.Load(".static-hoster.env")
	if err != nil {
		log.Println("Error loading .env file")
	} else {
		godotenv.Overload()
	}
	defaultApiKey := "test123"
	env = new(Environment)
	env.RootDir = envReader("STATIC_HOSTER_HOME", fmt.Sprintf("%s/static-hoster/", os.Getenv("HOME")))
	env.StaticDir = envReader("STATIC_HOSTER_HOST_DIR", fmt.Sprintf("%shosted/", env.RootDir))
	env.Port = fmt.Sprintf(":%s", envReader("STATIC_HOSTER_PORT", "8080"))
	env.User = envReader("STATIC_HOSTER_USER", "user")
	env.ApiKey = envReader("STATIC_HOSTER_API_KEY", defaultApiKey)
	env.BaseRoute = envReader("STATIC_HOSTER_BASE_ROUTE", "/home")
	if env.ApiKey == defaultApiKey {
		fmt.Printf("[STATIC-HOSTER-Warning]\t Environment Variable \"STATIC_HOSTER_API_KEY\" not set. Use default key \"%s\". DONT USE THIS FOR PRODUCTION!\n", defaultApiKey)
	}
	mkdirIfNotExist(env.RootDir, os.ModePerm)
	mkdirIfNotExist(env.StaticDir, os.ModePerm)
	return
}

func envReader(envVar string, defaultVal string) string {
	if os.Getenv(envVar) == "" {
		return defaultVal
	} else {
		return os.Getenv(envVar)
	}
}

func mkdirIfNotExist(dir string, perm fs.FileMode) {
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("[STATIC_HOSTER-Info]\tFolder %s does not exit, trying to create it\n", dir)
		err := os.Mkdir(dir, perm)
		if err != nil {
			log.Println(err)
		}
	}
}
