package config

import (
	"github.com/joho/godotenv"
	"os"
	"fmt"
	"strconv"
	"time"
)

func GetEnvStr(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetEnvInt(key string) (int, error ){
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 0, err
	}
	return value, nil
}

func GetEnvTime(key string) (time.Duration, error) {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}
	value, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		return 0, err
	}
	return value, nil
}