package config

import (
	"flag"
	"os"
)

var addr string
var baseURL string
var logLevel string
var fileStoragePath string

func ParseFlags() {
	flag.StringVar(&addr, "a", ":8080", "server address")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "base url")
	flag.StringVar(&logLevel, "l", "info", "log level")
	flag.StringVar(&fileStoragePath, "f", "default.json", "file storage path")
	flag.Parse()

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		addr = envRunAddr
	}

	if envBaseUrl := os.Getenv("BASE_URL"); envBaseUrl != "" {
		baseURL = envBaseUrl
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		logLevel = envLogLevel
	}

	if evnFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); evnFileStoragePath != "" {
		fileStoragePath = evnFileStoragePath
	}
}

func Addr() string {
	return addr
}

func BaseURL() string {
	return baseURL
}

func LogLevel() string {
	return logLevel
}

func FileStoragePath() string {
	return fileStoragePath
}
