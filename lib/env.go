package lib

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
    err := godotenv.Load(".env")
    if err != nil {
        return ""
    }
    values := os.Getenv(key)
    return values
}
    
