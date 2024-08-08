package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("NO .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HHTP_PORT", ":8080"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "sustainability_impact_service"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "03212164"))

	return config
}

func coalesce(key string, defaultVaule interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultVaule
}
