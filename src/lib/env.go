package lib

import (
	"fmt"
	"os"
)

func checkEnv(envName string) (string, error) {
	v := os.Getenv(envName)
	if v == "" {
		return "", fmt.Errorf(`env variable "%s" is not defined`, envName)
	}
	return v, nil
}

func GetDatabaseURL() string {
	url, err := checkEnv("DB_URL")
	Must(err)
	return url
}

func GetTestDatabaseURL() string {
	url, err := checkEnv("DB_TEST_URL")
	if err != nil {
		url = "postgres://postgres:example@127.0.0.1:5433/test_db?sslmode=disable"
	}
	return url
}

func GetRedisURL() string {
	url, err := checkEnv("REDIS_URL")
	Must(err)
	return url
}

func GetTokenSecret() string {
	secret, err := checkEnv("TOKEN_SECRET")
	Must(err)
	return secret
}

func GetServerBindAddress() string {
	host, err := checkEnv("HOST")
	if err != nil {
		host = "0.0.0.0"
	}
	port, err := checkEnv("PORT")
	if err != nil {
		port = "9091"
	}
	return fmt.Sprintf("%s:%s", host, port)
}

func GetServerReadTimeout() string {
	v, _ := checkEnv("SERVER_READ_TIMEOUT")
	if v == "" {
		return "60"
	}

	return v
}

func GetDevelopmentMode() string {
	v := os.Getenv("DEVELOPMENT_MODE")
	if v == "" {
		return "development"
	}
	return v
}
