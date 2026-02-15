package main

import (
	"os"

	"github.com/joho/godotenv"
)

type getIP struct {
	IPAddress string `json:"ipAddress"`
	MaxAgeInDays int `json:"maxAgeInDays"`
}

type headers

func main() {
	godotenv.Load()
	apiURL := os.Getenv("API_URL")
	apiKey := os.Getenv("ABUSEIP_KEY")
}
