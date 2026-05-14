package config

import (
	"flag"
	"os"
)

var addr string
var baseURL string

func ParseFlags() {
	flag.StringVar(&addr, "a", ":8080", "server address")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "base url")
	flag.Parse()

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		addr = envRunAddr
	}

	if envBaseUrl := os.Getenv("BASE_URL"); envBaseUrl != "" {
		baseURL = envBaseUrl
	}
}

func Addr() string {
	return addr
}

func BaseURL() string {
	return baseURL
}
