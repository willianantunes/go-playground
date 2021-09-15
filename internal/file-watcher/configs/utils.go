package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getEnvValueOrUseDefault(envName string, envDefault string) string {
	envValue, isEnvSet := os.LookupEnv(envName)

	if isEnvSet == false {
		return envDefault
	}

	return envValue
}

func getEnvValueOrPanic(envName string) string {
	envValue, isEnvSet := os.LookupEnv(envName)

	if isEnvSet == false {
		errorMessage := fmt.Sprintf("Must have variable is not set: %s", envName)
		panic(errorMessage)
	}

	return envValue
}

func retrieveFullPathFromRelative(relativePath string) string {
	return filepath.Join(directoryWhereAmI(), relativePath)
}

func directoryWhereAmI() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}
