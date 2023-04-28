package environment

import (
	"github.com/joho/godotenv"
	"os"
)

func Load() error {
	env := os.Getenv("APP_ENV")
	var envFile string

	switch env {
	case "test":
		envFile = ".env.test"
	default:
		envFile = ".env"
	}

	return godotenv.Load(envFile)
}
