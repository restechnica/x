package exec

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/restechnica/x/pkg/cli"
	"github.com/restechnica/x/pkg/cli/commands"
)

// Run will execute the CLI root command.
func Run() (err error) {
	var command = commands.NewRootCommand()

	if err = command.Execute(); err != nil {
		if errors.As(err, &cli.CommandError{}) {
			log.Error().Err(err).Msg("")
		}

		os.Exit(1)
	}

	return err
}
