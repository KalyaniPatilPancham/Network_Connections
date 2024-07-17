package fetcher

import (
	"demo/models"
	"encoding/json"
	"net/http"
)

// Fetcher defines the interface for fetching JSON data
type Fetcher interface {
	FetchJSONData(url string) (*models.Data, error)
}

// defaultFetcher implements Fetcher
type defaultFetcher struct{}

// FetchJSONData implements the Fetcher interface for defaultFetcher
func (f *defaultFetcher) FetchJSONData(url string) (*models.Data, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data models.Data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

// NewFetcher returns a new instance of Fetcher
func NewFetcher() Fetcher {
	return &defaultFetcher{}
}
