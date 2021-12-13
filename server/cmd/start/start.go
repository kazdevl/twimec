package start

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start <twitter_token>",
		Short: "start proccessings with twitter_token",
		Long:  `start is for getting twitter_image_contents and providing a well-formatted UI`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			proccess(args[0])
		},
	}
}

func proccess(token string) {
	// TODO impl
}
