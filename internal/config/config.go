package config

import "flag"

var addr string
var baseURL string

func ParseFlags() {
	flag.StringVar(&addr, "a", ":8080", "server address")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "base url")
	flag.Parse()
}

func Addr() string {
	return addr
}

func BaseURL() string {
	return baseURL
}
