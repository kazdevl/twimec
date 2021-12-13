package start

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/kazdevl/twimec/api"
	"github.com/kazdevl/twimec/repository/local"
	"github.com/kazdevl/twimec/usecase"
	"github.com/labstack/echo/v4"
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
	// 1. launch api server
	e := echo.New()
	api.RegisterRoutes(e)
	go func() {
		if err := e.Start(":6666"); err != nil {
			log.Fatal(err)
		}
	}()

	// 2. call cron job
	tclient := usecase.NewTClient(false, token)
	homeDir, _ := os.UserHomeDir()
	chapterRepo := local.NewChapterRepository(fmt.Sprintf("%s/twimec/storage/contents", homeDir))
	configRepo := local.NewConfigRepository(fmt.Sprintf("%s/twimec/storage/config/contents", homeDir))
	cronjob := usecase.NewCronjob(tclient, chapterRepo, configRepo)

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("21:00").Do(cronjob.FetchContents)
	s.StartAsync()

	// 3. Graceful shatdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
