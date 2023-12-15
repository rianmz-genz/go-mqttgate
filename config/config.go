package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + key)
	}
	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")
	print(s3Bucket)
	print(secretKey)
	return os.Getenv(key)
}
