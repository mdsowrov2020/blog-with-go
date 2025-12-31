package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  bool
}

type Config struct {
	Version      string
	ServiceName  string
	JWTSecretKey string
	HTTPPort     int
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env file", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT key  is required")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP Port  is required")
		os.Exit(1)
	}
	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("Failed to convert port")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("Host  is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("PORT  is required")
		os.Exit(1)
	}

	dbPortINT, err := strconv.Atoi(dbPort)
	if err != nil {
		fmt.Println("Failed to convert port")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name  is required")
		os.Exit(1)
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("User  is required")
		os.Exit(1)
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("Password  is required")
		os.Exit(1)
	}
	sslMode := os.Getenv("DB_SSLMODE")
	sslModeParsed, err := strconv.ParseBool(sslMode)
	if err != nil {
		fmt.Println("Invalid ssl mode value")
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		DBHost:     dbHost,
		DBPort:     dbPortINT,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBSSLMode:  sslModeParsed,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		JWTSecretKey: jwtSecretKey,
		HTTPPort:     port,
		DB:           dbConfig,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}

	return configurations
}
