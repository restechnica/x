package v1

import (
	"github.com/spf13/cobra"
)

// NewGetCommand creates a new get command.
// Returns a new init spf13/cobra command.
func NewGetCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "get",
	}

	return command
}
