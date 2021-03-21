package main

import (
	todo "github.com/RymarSergey/my_todo"
	"github.com/RymarSergey/my_todo/pkg/handler"
	"log"
)

func main() {
	handlers:=new(handler.Handler)
	srv:=new(todo.Server)
	if err:=srv.Run("8000",handlers.InitRoutes());err!=nil{
		log.Fatalf("err ocured while runing http server: %s",err.Error())
	}
}
