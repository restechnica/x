package commands

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/restechnica/x/pkg/cli"
	"github.com/restechnica/x/pkg/cli/commands/v1"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
}

// NewRootCommand creates a new root command.
// Returns the new spf13/cobra command.
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:               "x",
		PersistentPreRunE: RootCommandPersistentPreRunE,
	}

	command.PersistentFlags().BoolVarP(&cli.VerboseFlag, "verbose", "v", false, "increase log level verbosity to Info")
	command.PersistentFlags().BoolVarP(&cli.DebugFlag, "debug", "d", false, "increase log level verbosity to Debug")

	command.AddCommand(v1.NewV1Command())
	command.AddCommand(v1.NewGetCommand())

	return command
}

// RootCommandPersistentPreRunE runs before the command and any subcommand runs.
// Returns an error if it failed.
func RootCommandPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	// silence usage and errors because errors at this point are unrelated to CLI usage errors
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true

	ConfigureLogging()

	return err
}

func ConfigureLogging() {
	SetLogLevel()
}

func SetLogLevel() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	if cli.VerboseFlag {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if cli.DebugFlag {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
