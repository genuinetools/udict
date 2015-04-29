package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIResponse struct {
	Results []Result `json:"list"`
	Tags    []string `json:"tags"`
	Type    string   `json:"result_type"`
}

type Result struct {
	Id         int32  `json:"defid"`
	Author     string `json:"author"`
	Definition string `json:"definition"`
	Link       string `json:"permalink"`
	ThumbsDown int32  `json:"thumbs_down"`
	ThumbsUp   int32  `json:"thumbs_up"`
	Word       string `json:"word"`
}

func DefineWord(word string) (response *APIResponse, err error) {

	endpoint := fmt.Sprintf("http://api.urbandictionary.com/v0/define?page=%d&term=%s", 1, word)
	resp, err := http.Get(endpoint)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	if err != nil {
		return nil, err
	}

	return
}
