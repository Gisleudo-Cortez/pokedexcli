package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2/location-area/?limit=20"

type LocationAreaResponse struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous,omitempty"` // This is used for optional fields
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationAreaURL(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []string{}, err
	}

	var locationAreaResp LocationAreaResponse
	err = json.Unmarshal(body, &locationAreaResp)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return []string{}, err
	}

	var out []string

	// Iterate over the results and add each name to the list
	for _, result := range locationAreaResp.Results {
		out = append(out, result.Name)
	}
	return out, nil
}
