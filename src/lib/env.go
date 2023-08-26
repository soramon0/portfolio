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
	Must(err)
	port, err := checkEnv("PORT")
	Must(err)

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
