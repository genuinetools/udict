package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	VERSION = "v0.1.0"
	BANNER  = `           _ _      _
 _   _  __| (_) ___| |_
| | | |/ _` + "`" + ` | |/ __| __|
| |_| | (_| | | (__| |_
 \__,_|\__,_|_|\___|\__|

 Urban Dictionary Command Line Tool
 Version: ` + VERSION
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

func main() {
	var version bool

	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.Parse()

	args := os.Args

	if len(args) <= 1 || args[1] == "help" {
		fmt.Println(BANNER)
		return
	}

	if version || args[1] == "version" {
		fmt.Println(VERSION)
		return
	}

	word := args[1]
	endpoint := fmt.Sprintf("http://api.urbandictionary.com/v0/define?page=%d&term=%s", 1, word)
	resp, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Request to %q failed: %v", endpoint, err)
		return
	}
	defer resp.Body.Close()

	/*
		Reponse comes back like:
		{
		  "tags": [
		    "icbw",
		    "+1",
		    "fine"
		  ],
		  "result_type": "exact",
		  "list": [
		    {
		      "defid": 2535452,
		      "word": "LGTM",
		      "author": "akshay_s",
		      "permalink": "http://lgtm.urbanup.com/2535452",
		      "definition": "An acronym for \"Looks Good To Me\", often used as a quick response after reviewing someones essay, code, or design document.,example:[LGTM], dude. You can go ahead and push this [craxy] code to the prod server. Well make [M$] wish they were flippin burgers!  [Woot]!\r\n",
		      "thumbs_up": 256,
		      "thumbs_down": 39,
		      "current_vote": ""
		    },
		    ...
		  ],
		  "sounds": []
		}
	*/

	var response APIResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	if err != nil {
		fmt.Printf("Decoding api response as JSON failed: %v", err)
		return
	}

	defResponse := fmt.Sprintf("%d definitions returned\n", len(response.Results))

	for _, def := range response.Results {
		defResponse += fmt.Sprintf("\n%s\n--(%s) <%s>\n", def.Definition, def.Word, def.Link)
	}

	fmt.Println(defResponse)
}
