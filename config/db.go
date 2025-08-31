package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ctxKey struct{}
var DBKey = ctxKey{}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

