package main

import (
	"log"

	authorization "github.com/Domitory66/AuthGo"
	"github.com/Domitory66/AuthGo/pkg/handler"
	"github.com/Domitory66/AuthGo/pkg/repository"
	"github.com/Domitory66/AuthGo/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("error initialization configs %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(authorization.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
