package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config структура для хранения конфигурационных данных
type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	LogLevel   string
	Message    string
}

// LoadConfig загружает конфигурацию из .env файла, переменных окружения и config.json
func LoadConfig() Config {
	// Загрузка .env файла
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading configuration from environment variables")
	}

	// Чтение переменных окружения
	config := Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "hello_api_db"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
	}

	// Загрузка config.json
	loadMessageConfig(&config)

	return config
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// loadMessageConfig загружает текст сообщения из config.json
func loadMessageConfig(config *Config) {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Printf("Failed to open config.json: %v", err)
		config.Message = "Default message"
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var jsonData struct {
		Message string `json:"message"`
	}
	if err := decoder.Decode(&jsonData); err != nil {
		log.Printf("Failed to parse config.json: %v", err)
		config.Message = "Default message"
		return
	}

	config.Message = jsonData.Message
}
