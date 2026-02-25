package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	ipArg := os.Args[1:]
	ip := ipArg[0]
	godotenv.Load()
	apiKey := os.Getenv("ABUSEIP_KEY")

	params := map[string]string{
		"ipAddress":    ip,
		"maxAgeInDays": "90",
	}

	cfg, err := newConfig(apiKey, "check", params)
	if err != nil {
		log.Fatalf("failed to create api config: %s", err)
	}

	body, err := cfg.abusedClient.FetchIPData(cfg.fullURL, cfg.apiKey)
	if err != nil {
		log.Fatalf("%s", err)
	}

	var prettyPrint bytes.Buffer
	if err := json.Indent(&prettyPrint, body, "", "		"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(prettyPrint.String())
}
