package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	urbanDictionaryAPIURI = "https://api.urbandictionary.com/v0"
)

// Response describes the data structure that comes back from the
// Urban Dictionary API.
type Response struct {
	Results []Result `json:"list"`
	Tags    []string `json:"tags"`
	Type    string   `json:"result_type"`
}

// Result holds the information for a given definition.
type Result struct {
	ID         int64  `json:"defid"`
	Author     string `json:"author"`
	Definition string `json:"definition"`
	Link       string `json:"permalink"`
	ThumbsDown int64  `json:"thumbs_down"`
	ThumbsUp   int64  `json:"thumbs_up"`
	Word       string `json:"word"`
}

// Define returns the definitions from Urban Dictionary for a given word.
func Define(word string) (response *Response, err error) {
	endpoint := fmt.Sprintf("%s/define?page=%d&term=%s", urbanDictionaryAPIURI, 1, url.QueryEscape(word))
	resp, err := http.Get(endpoint)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
