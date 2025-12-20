package config

import (
	"fmt"
	"os"
	"strconv"
)

var configurations *Config

type Config struct {
	Version      string
	ServiceName  string
	JWTSecretKey string
	HTTPPort     int
}

func loadConfig() {
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
	}
	jwtSecretKey := os.Getenv("JWT_SECRET")
	if jwtSecretKey == "" {
		fmt.Println("JWT key  is required")
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP Port  is required")
	}
	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("Failed to convert port")
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		JWTSecretKey: jwtSecretKey,
		HTTPPort:     port,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}

	return configurations
}
