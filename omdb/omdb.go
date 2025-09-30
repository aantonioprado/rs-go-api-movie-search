package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Result struct {
	Search       []SearchResult `json:"Search"`
	TotalResults string         `json:"totalResults"`
	Response     string         `json:"Response"`
}

type SearchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func Search(apiKey, title string) (Result, error) {
	v := url.Values{}

	v.Set("apikey", apiKey)
	v.Set("s", title)

	res, err := http.Get("https://www.omdbapi.com/?" + v.Encode())

	if err != nil {
		return Result{}, fmt.Errorf("faield to make request to omdb: %w", err)
	}

	defer res.Body.Close()

	var result Result

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return Result{}, fmt.Errorf("faield to make response to omdb: %w", err)
	}

	return result, nil
}
