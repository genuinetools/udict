package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/genuinetools/pkg/cli"
	"github.com/genuinetools/udict/api"
	"github.com/genuinetools/udict/version"
	"github.com/sirupsen/logrus"
)

var (
	debug bool
)

func main() {
	// Create a new cli program.
	p := cli.NewProgram()
	p.Name = "udict"
	p.Description = "A command line urban dictionary"

	// Set the GitCommit and Version.
	p.GitCommit = version.GITCOMMIT
	p.Version = version.VERSION

	// Setup the global flags.
	p.FlagSet = flag.NewFlagSet("global", flag.ExitOnError)
	p.FlagSet.BoolVar(&debug, "d", false, "enable debug logging")

	// Set the before function.
	p.Before = func(ctx context.Context) error {
		// Set the log level.
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		if p.FlagSet.NArg() < 1 {
			return errors.New("pass a word or phrase")
		}

		return nil
	}

	// Set the main program action.
	p.Action = func(ctx context.Context, args []string) error {
		// On ^C, or SIGTERM handle exit.
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)
		go func() {
			for sig := range c {
				logrus.Infof("Received %s, exiting.", sig.String())
				os.Exit(0)
			}
		}()

		word := strings.Join(args, " ")

		response, err := api.Define(word)
		if err != nil {
			return fmt.Errorf("decoding API response failed: %v", err)
		}

		defResponse := fmt.Sprintf("%d definitions returned\n", len(response.Results))

		for _, def := range response.Results {
			defResponse += fmt.Sprintf("\n%s\n--(%s) <%s>\n", def.Definition, def.Word, def.Link)
		}

		fmt.Println(defResponse)

		return nil
	}

	// Run our program.
	p.Run()
}
