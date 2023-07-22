package exec

import (
	"errors"
	"log"
	"os"

	"github.com/restechnica/x/pkg/cli"
	"github.com/restechnica/x/pkg/cli/commands"
)

// Run will execute the CLI root command.
func Run() (err error) {
	var command = commands.NewRootCommand()

	if err = command.Execute(); err != nil {
		if errors.As(err, &cli.CommandError{}) {
			log.Fatal(err)
		}

		os.Exit(1)
	}

	return err
}
