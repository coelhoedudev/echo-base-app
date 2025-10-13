package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type ServerConfig struct {
	JWTSecret   string
	JWTDuration string
	Port        string
	RedisUrl    string
}

var (
	once   sync.Once
	config *Config
)

func LoadConfig() {
	godotenv.Load()

	once.Do(func() {
		config = &Config{
			DB: DBConfig{
				Host:     getEnv("DB_HOST", "default"),
				Port:     getEnv("DB_PORT", "default"),
				Name:     getEnv("DB_NAME", "default"),
				User:     getEnv("DB_USER", "default"),
				Password: getEnv("DB_PASSWORD", "default"),
			},
			Server: ServerConfig{
				Port:        getEnv("SERVER_PORT", "default"),
				JWTSecret:   getEnv("JWT_SECRET", "default"),
				JWTDuration: getEnv("JWT_EXPIRES_IN", "default"),
				RedisUrl:    getEnv("REDIS_URL", "default"),
			},
		}
	})
}

func getEnv(key string, fallback string) string {
	env, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return env
}

func GetConfig() *Config {
	return config
}
