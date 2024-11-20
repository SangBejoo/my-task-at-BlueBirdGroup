package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    Tile38Host string
    Tile38Port string
}

func LoadConfig() *Config {
    // Load environment variables from .env file if it exists
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    config := &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "root"),
        DBPassword: getEnv("DB_PASSWORD", "secret"),
        DBName:     getEnv("DB_NAME", "subagiyaa"),
        Tile38Host: getEnv("TILE38_HOST", "localhost"),
        Tile38Port: getEnv("TILE38_PORT", "9851"),
    }

    return config
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}