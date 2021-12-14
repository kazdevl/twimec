package stop

import (
	"os"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "stop processings",
		Long:  `stop is for stopping processings`,
		Run: func(cmd *cobra.Command, args []string) {
			os.Exit(0)
		},
	}
	return cmd
}
