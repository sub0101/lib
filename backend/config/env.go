package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}
func GetEnv(s string) string {
	return os.Getenv(s)
}
