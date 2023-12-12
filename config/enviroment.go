package config

import (
	"log"

	"github.com/joho/godotenv"
)

var envs = map[string]string{}

func initEnv(envFile string, envs *map[string]string) {
	var err error
	var env map[string]string
	if envFile == "dev" {
		env, err = godotenv.Read(".env.development")
		if err != nil {
			log.Fatal("Error loading .env.development file")
		}
	} else if envFile == "prod" {
			env, err = godotenv.Read(".env.production")
			if err != nil {
				log.Fatal("Error loading .env.production file")
			}
	}

	for key, value := range env {
		(*envs)[key] = value
	}

}

func GetEnv(key string) string {
	if len(envs) == 0 {
		initEnv("dev", &envs)
	}
	return envs[key]
}

var PORT = GetEnv("PORT")