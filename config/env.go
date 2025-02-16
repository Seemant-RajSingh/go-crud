package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv" // package to directly read from .env files
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "SQL24@MyDBSrS"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "goecom"),       // schema not generated on startup
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7), // getEnv("JWT_EXP", 3600*24*7) dosent take int
		JWTSecret:              getEnv("JWT_SECRET", "not-secret-secret-anymore?"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		// convert time from string to int64 (MUST STORE JWTExiration in string in .env)
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
