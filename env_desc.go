package helpers

import "fmt"

func EnvDesc(description, envVar string) string {
	return fmt.Sprintf("%s Can be set with the %s env variable", description, envVar)
}
