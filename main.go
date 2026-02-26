package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/CodeZeroSugar/ip-guard/internal/process"
	"github.com/CodeZeroSugar/ip-guard/internal/scanner"
	"github.com/joho/godotenv"
)

func main() {
	scanner := scanner.IPScanner{Runner: process.OSRunner{}}
	scanResults, err := scanner.GetRemoteIPs()
	if err != nil {
		log.Fatalf("scanner failed to obtain remote connections: %s", err)
	}

	godotenv.Load()
	apiKey := os.Getenv("ABUSEIP_KEY")
	params := make(map[string]string)

	for ip := range scanResults {
	}

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
