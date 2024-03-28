package helpers

import (
	"os"
	"strings"
)

const LocalEnv = "local"

func GetEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = LocalEnv
	}
	return strings.ToLower(env)
}
