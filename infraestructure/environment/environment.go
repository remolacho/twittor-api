package environment

import (
	"errors"
	"github.com/joho/godotenv"
)

func Load(opts ...string) error {
	if opts == nil {
		return godotenv.Load(".env")
	}

	if opts[0] == "test" {
		return godotenv.Load("../../../.env.test")
	}

	return errors.New("no implement env fila")
}
