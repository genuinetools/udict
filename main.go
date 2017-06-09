package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jessfraz/udict/api"
	"github.com/jessfraz/udict/version"
)

const (
	// BANNER is what is printed for help/info output.
	BANNER = `           _ _      _
 _   _  __| (_) ___| |_
| | | |/ _` + "`" + ` | |/ __| __|
| |_| | (_| | | (__| |_
 \__,_|\__,_|_|\___|\__|

 Urban Dictionary Command Line Tool
 Version: %s

`
)

var (
	vrsn bool
)

func init() {
	// parse flags
	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&vrsn, "v", false, "print version and exit (shorthand)")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, version.VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("udict version %s, build %s", version.VERSION, version.GITCOMMIT)
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		usageAndExit("Pass a word or phrase.", 1)
	}

	// parse the arg
	arg := flag.Args()[0]

	if arg == "help" {
		usageAndExit("", 0)
	}

	if arg == "version" {
		fmt.Printf("udict version %s, build %s", version.VERSION, version.GITCOMMIT)
		os.Exit(0)
	}
}

func main() {
	word := flag.Args()[0]

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

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}
