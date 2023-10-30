package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func initEnv(env string) {
	if env == "dev" {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env.development file")
	}
	} else if env == "prod" {
		err := godotenv.Load(".env.production")
		if err != nil {
			log.Fatal("Error loading .env.production file")
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

}

var envs = map[string]string{}

func loadEnv (envs *map[string]string) {
	var enviroment = os.Getenv("ENV")
	initEnv(enviroment)
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	for key, value := range env {
		(*envs)[key] = value
	}
}

func GetEnv(key string) string {
	loadEnv(&envs)
	return envs[key]
}




var PORT = GetEnv("PORT")