package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/kazdevl/twimec/api"
	"github.com/kazdevl/twimec/usecase"
	"github.com/labstack/echo/v4"
)

func init() {
	usecase.GetCredentialH = make(chan struct{}, 1)
}

func main() {
	// 1. launch api server
	e := echo.New()
	api.RegisterRoutes(e)
	go func() {
		if err := e.Start(":8080"); err != nil {
			log.Fatal(err)
		}
	}()

	// 2. wait until server get credential
	<-usecase.GetCredentialH
	close(usecase.GetCredentialH)

	// 3. call cron job
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("21:00").Do(usecase.FetchContents)
	s.StartAsync()

	// 4. Graceful shatdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
