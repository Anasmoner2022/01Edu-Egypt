package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

// FetchData gets data from a URL and unmarshals into target
func FetchData(url string, target interface{}) error {
	// 1. Make HTTP GET request using http.Get()
	// 2. Check for errors
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Can't Get URL: %v", err)
	}
	// 3. Defer close response body
	defer resp.Body.Close()
	// 4. Check status code (should be 200)
	if resp.StatusCode != 200 {
		log.Fatal("Wrong Status codes")
	}
	// 5. Use json.NewDecoder to decode response into target
	err = json.NewDecoder(resp.Body).Decode(target)
	// 6. Return any errors
	return err
}
