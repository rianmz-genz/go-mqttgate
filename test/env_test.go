package test

import (
	"adriandidimqttgate/config"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	assert.NotNil(t, s3Bucket)
	assert.NotNil(t, secretKey)

	println(s3Bucket)
	println(secretKey)
}

func TestGetEnv(t *testing.T) {
	username := config.GetEnv("DB_USERNAME")

	assert.Equal(t, "super", username)
}
