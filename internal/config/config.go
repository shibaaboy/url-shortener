package config

import (
	"flag"
	"os"
)

var addr string
var baseURL string
var logLevel string

func ParseFlags() {
	flag.StringVar(&addr, "a", ":8080", "server address")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "base url")
	flag.StringVar(&logLevel, "l", "info", "log level")
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
