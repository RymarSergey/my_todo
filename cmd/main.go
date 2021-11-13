package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	todo "github.com/RymarSergey/my_todo"
	"github.com/RymarSergey/my_todo/pkg/handler"
	"github.com/RymarSergey/my_todo/pkg/repository"
	"github.com/RymarSergey/my_todo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialazing config :%s ", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loaing env variables:%s ", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize DB :%s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("err ocured while runing http server: %s", err.Error())
		}
	}()
	logrus.Info("App start ...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("App stop")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error ocured on server shuting down:%s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Error ocured on db connection close:%s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
