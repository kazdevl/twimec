package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kazdevl/twimec/cmd/twimec/set"
	"github.com/kazdevl/twimec/cmd/twimec/start"
	"github.com/kazdevl/twimec/cmd/twimec/stop"
	"github.com/kazdevl/twimec/repository/local"
	"github.com/spf13/cobra"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	var (
		contentsPath       = fmt.Sprintf("%s/twimec/storage/contents", homeDir)
		contentsConfigPath = fmt.Sprintf("%s/twimec/storage/config/contents", homeDir)
		templatePath       = fmt.Sprintf("%s/twimec/web/template", homeDir)
	)
	if err := os.MkdirAll(contentsPath, 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(contentsConfigPath, 0777); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(templatePath, 0777); err != nil {
		log.Fatal(err)
	}
	tDir, err := filepath.Abs("./../../web/template")
	if err != nil {
		log.Fatal(err)
	}
	des, err := os.ReadDir(tDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, de := range des {
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", tDir, de.Name()))
		if err != nil {
			log.Fatal(err)
		}
		dest, err := os.Create(fmt.Sprintf("%s/%s", templatePath, de.Name()))
		defer dest.Close()
		if err != nil {
			log.Fatal(err)
		}
		dest.Write(data)
	}

	chapterRepo := local.NewChapterRepository(contentsPath)
	configRepo := local.NewConfigRepository(contentsConfigPath)
	contentinfoRepo := local.NewContentInfoRepository(contentsConfigPath)

	var rootCmd = &cobra.Command{Use: "twimec"}
	rootCmd.AddCommand(
		start.NewCmd(configRepo, contentinfoRepo, chapterRepo),
		set.NewCmd(configRepo, contentinfoRepo),
		stop.NewCmd(),
	)
	cobra.CheckErr(rootCmd.Execute())
}
