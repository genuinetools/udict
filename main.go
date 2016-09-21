package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jfrazelle/udict/api"
)

const (
	// VERSION is the binary version.
	VERSION = "v0.2.0"
	// BANNER is what is printed for help/info output.
	BANNER = `           _ _      _
 _   _  __| (_) ___| |_
| | | |/ _` + "`" + ` | |/ __| __|
| |_| | (_| | | (__| |_
 \__,_|\__,_|_|\___|\__|

 Urban Dictionary Command Line Tool
 Version: ` + VERSION
)

var (
	version bool
)

func init() {
	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.Parse()
}

func main() {
	// parse the args
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

	response, err := api.Define(word)
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
