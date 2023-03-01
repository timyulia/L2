package main

import (
	"calendar"
	"calendar/pkg/handler"
	"calendar/pkg/repository"
	"context"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	handlers := handler.NewHandler(repos)

	srv := new(calendar.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.EventLogger(handlers.InitRoutes())); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Print("started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
