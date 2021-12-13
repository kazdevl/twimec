package main

import (
	"fmt"
	"os"

	"github.com/kazdevl/twimec/cmd/set"
	"github.com/kazdevl/twimec/cmd/start"
	"github.com/kazdevl/twimec/cmd/stop"
	"github.com/kazdevl/twimec/repository/local"
	"github.com/spf13/cobra"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	const contentsPath = "twimec/storage/contents"
	const contentsConfigsPath = "twimec/storage/config/contents"
	chapterRepo := local.NewChapterRepository(fmt.Sprintf("%s/%s", homeDir, contentsPath))
	configRepo := local.NewConfigRepository(fmt.Sprintf("%s/%s", homeDir, contentsConfigsPath))
	contentinfoRepo := local.NewContentInfoRepository(fmt.Sprintf("%s/%s", homeDir, contentsConfigsPath))

	var rootCmd = &cobra.Command{Use: "twimec"}
	rootCmd.AddCommand(
		start.NewCmd(configRepo, chapterRepo),
		set.NewCmd(configRepo, contentinfoRepo),
		stop.NewCmd(),
	)
	cobra.CheckErr(rootCmd.Execute())
}
