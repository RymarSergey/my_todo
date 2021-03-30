package main

import (
	todo "github.com/RymarSergey/my_todo"
	"github.com/RymarSergey/my_todo/pkg/handler"
	"github.com/RymarSergey/my_todo/pkg/repository"
	"github.com/RymarSergey/my_todo/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err:=initConfig();err != nil {
		log.Fatalf("error initialazing config :%s ",err.Error())
	}
	repo:=repository.NewRepository()
	services:=service.NewService(repo)
	handlers:=handler.NewHandler(services)

	srv:=new(todo.Server)
	if err:=srv.Run(viper.GetString("port"),handlers.InitRoutes());err!=nil{
		log.Fatalf("err ocured while runing http server: %s",err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
