package utility

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnv(key string, filePath string) string {
	if err := godotenv.Load(filePath); err != nil {
		log.Println("failed to open .env file")
	}

	return os.Getenv(key)
}
