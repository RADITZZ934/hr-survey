package config

import "os"

type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	JWTSecret string
}

var AppConfig Config

func LoadConfig() {
	AppConfig = Config{
		Port: os.Getenv("1302"),

		DBHost:     os.Getenv("localhost"),
		DBPort:     os.Getenv("5432"),
		DBUser:     os.Getenv("postgres"),
		DBPassword: os.Getenv("postgres"),
		DBName:     os.Getenv("hr_survey"),
		DBSSLMode:  os.Getenv("disable"),

		JWTSecret: os.Getenv("supersecret123"),
	}
}
