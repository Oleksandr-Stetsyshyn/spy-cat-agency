package services

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	validBreeds     []string
	lastFetchedTime time.Time
	cacheDuration   = time.Hour
	mu              sync.Mutex
)

type Breed struct {
	Name string `json:"name"`
}

func FetchValidBreeds() error {
	mu.Lock()
	defer mu.Unlock()

	if time.Since(lastFetchedTime) < cacheDuration {
		return nil
	}

	log.Println("Fetching valid breeds")
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var breeds []Breed

	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		return err
	}

	validBreeds = nil
	for _, breed := range breeds {
		validBreeds = append(validBreeds, breed.Name)
	}

	lastFetchedTime = time.Now()
	return nil
}

func IsValidBreed(breed string) bool {
	if len(validBreeds) == 0 {
		if err := FetchValidBreeds(); err != nil {
			return false
		}
	}
	for _, b := range validBreeds {
		if b == breed {
			return true
		}
	}
	return false
}
