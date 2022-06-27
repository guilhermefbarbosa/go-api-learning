package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Config struct {
	AppName string
	AppEnv  string
	Tokens  []string
}

var globalConfig Config

func init() {
	var config Config

	envs, err := godotenv.Read()
	if err != nil {
		fmt.Printf("Could not load dotenv: %s. Loading environment variables \n", err)
	}

	config.AppEnv = loadEnvVariable("ENV", envs)

	var keysList []string
	keysList = append(keysList, loadEnvVariable("MAIN_TOKEN", envs))
	keysList = append(keysList, loadEnvVariable("ANOTHER_TOKEN", envs))

	config.Tokens = keysList

	if isTestEnvironment(config.AppEnv) {
		fmt.Println("running test environment")
	}
	globalConfig = config
}

func loadEnvVariable(str string, envs map[string]string) string {
	ret := os.Getenv(str)
	if len(ret) > 0 {
		return ret
	}
	return envs[str]
}

func GetConfig() Config {
	return globalConfig
}

func isTestEnvironment(appEnv string) bool {
	return strings.HasSuffix(os.Args[0], ".test") || appEnv == "mock"
}
