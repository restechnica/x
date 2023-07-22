package v1

import (
	"github.com/spf13/cobra"
)

// NewListCommand creates a new list command.
// Returns a new list spf13/cobra command.
func NewListCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "list",
		RunE:  ListCommandRunE,
		Short: "Lists all available targets",
	}

	return command
}

// ListCommandRunE runs the list command.
// Returns an error if the command failed.
func ListCommandRunE(cmd *cobra.Command, args []string) (err error) {
	return err
}
