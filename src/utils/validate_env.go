package utils

import (
	"errors"
	"os"
	"strings"
)

func ValidateEnv(envs []string) error {
	not_sets := []string{}
	for _, env := range envs {
		if os.Getenv(env) == "" {
			not_sets = append(not_sets, env)
		}
	}
	if len(not_sets) > 0 {
		return errors.New("The following environment variables are not set: " + strings.Join(not_sets, ", "))
	}

	return nil
}
