package utils

import "os"

type RuntimeEnvironment string

const (
	RuntimeEnvironmentLocal        RuntimeEnvironment = "local"
	RuntimeEnvironmentGCPCloudRun  RuntimeEnvironment = "gcp_cloud_run"
	RuntimeEnvironmentAWSLambda    RuntimeEnvironment = "aws_lambda"
	RuntimeEnvironmentGCPAppEngine RuntimeEnvironment = "gcp_app_engine"
)

func IsValidRuntimeEnvironment(environment RuntimeEnvironment) bool {
	return environment == RuntimeEnvironmentLocal ||
		environment == RuntimeEnvironmentGCPCloudRun ||
		environment == RuntimeEnvironmentAWSLambda ||
		environment == RuntimeEnvironmentGCPAppEngine
}

func DetectRuntimeEnvironment() string {
	switch {
	case os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "":
		return string(RuntimeEnvironmentAWSLambda)
	case os.Getenv("K_SERVICE") != "":
		return string(RuntimeEnvironmentGCPCloudRun)
	case os.Getenv("GAE_ENV") == "standard":
		return string(RuntimeEnvironmentGCPAppEngine)
	default:
		return string(RuntimeEnvironmentLocal)
	}
}

func IsRuntimeEnvironmentGCPCloudRun() bool {
	return DetectRuntimeEnvironment() == string(RuntimeEnvironmentGCPCloudRun)
}

func IsRuntimeEnvironmentAWSLambda() bool {
	return DetectRuntimeEnvironment() == string(RuntimeEnvironmentAWSLambda)
}

func IsRuntimeEnvironmentGCPAppEngine() bool {
	return DetectRuntimeEnvironment() == string(RuntimeEnvironmentGCPAppEngine)
}

func IsRuntimeEnvironmentLocal() bool {
	return DetectRuntimeEnvironment() == string(RuntimeEnvironmentLocal)
}
