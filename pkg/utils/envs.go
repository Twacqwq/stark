package utils

import (
	"os"
	"strings"
)

func IsDev() bool {
	env := os.Getenv("ENV")
	return strings.Contains(strings.ToLower(env), "dev")
}
