package coinmate

import "fmt"

func EnvVarMissing(envVarName string) error {
	return fmt.Errorf("Missing enviroment variable: %s", envVarName)
}
