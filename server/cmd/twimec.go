package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kazdevl/twimec/cmd/set"
	"github.com/kazdevl/twimec/cmd/start"
	"github.com/kazdevl/twimec/cmd/stop"
	"github.com/kazdevl/twimec/repository/local"
	"github.com/spf13/cobra"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	var (
		contentsPath       = fmt.Sprintf("%s/twimec/storage/contents", homeDir)
		contentsConfigPath = fmt.Sprintf("%s/twimec/storage/config/contents", homeDir)
	)
	if err := os.MkdirAll(contentsPath, 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(contentsConfigPath, 0777); err != nil {
		log.Fatal(err)
	}

	chapterRepo := local.NewChapterRepository(contentsPath)
	configRepo := local.NewConfigRepository(contentsConfigPath)
	contentinfoRepo := local.NewContentInfoRepository(contentsConfigPath)

	var rootCmd = &cobra.Command{Use: "twimec"}
	rootCmd.AddCommand(
		start.NewCmd(configRepo, chapterRepo),
		set.NewCmd(configRepo, contentinfoRepo),
		stop.NewCmd(),
	)
	cobra.CheckErr(rootCmd.Execute())
}
