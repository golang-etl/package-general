package utils

import (
	"os"
	"strings"
)

func IsValidEnv() bool {
	return IsProductionEnv() || IsDevelopEnv() || IsStagingEnv() || IsLocalEnv()
}

func IsProductionEnv() bool {
	envVar := os.Getenv("ENV")
	lowercaseEnv := strings.ToLower(envVar)

	return lowercaseEnv == "production" || lowercaseEnv == "prod" || lowercaseEnv == ""
}

func IsDevelopEnv() bool {
	envVar := os.Getenv("ENV")
	lowercaseEnv := strings.ToLower(envVar)

	return lowercaseEnv == "develop" || lowercaseEnv == "dev"
}

func IsStagingEnv() bool {
	envVar := os.Getenv("ENV")
	lowercaseEnv := strings.ToLower(envVar)

	return lowercaseEnv == "staging" || lowercaseEnv == "stg"
}

func IsLocalEnv() bool {
	envVar := os.Getenv("ENV")
	lowercaseEnv := strings.ToLower(envVar)

	return lowercaseEnv == "local"
}
