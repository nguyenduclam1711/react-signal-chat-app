package env

import (
	"log"

	"github.com/joho/godotenv"
)

var EnvData map[string]string

func LoadEnvData() {
	loadedEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Cannot load env")
	}
	EnvData = loadedEnv
}
