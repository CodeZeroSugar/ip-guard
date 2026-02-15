package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const TEST_IP = "147.185.132.120"

func main() {
	godotenv.Load()
	apiURL := os.Getenv("API_URL")
	apiKey := os.Getenv("ABUSEIP_KEY")

	checkURL := apiURL + "/check"

	params := url.Values{}
	params.Add("ipAddress", TEST_IP)
	params.Add("maxAgeInDays", "90")

	fullURL := fmt.Sprintf("%s?%s", checkURL, params.Encode())

	client := &http.Client{}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var prettyPrint bytes.Buffer
	if err := json.Indent(&prettyPrint, body, "", "		"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(prettyPrint.String())
}
