package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {
	env := make(chan string, 1)

	godotenv.Load(".env")
	env <- os.Getenv(key)

	return <-env
}
