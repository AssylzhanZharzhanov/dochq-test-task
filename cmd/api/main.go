package main

import (
	"context"
	"github.com/AssylzhanZharzhanov/dochq-test-task/db/postgres"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/handler"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/service"
	"github.com/AssylzhanZharzhanov/dochq-test-task/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := new(server.Server)

	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error in reading env file: %s", err.Error())
	}
	port := viper.GetString("server.port")

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		DBName: viper.GetString("db.dbName"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Error in connecting to db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	go func() {
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error in starting server: %s", err.Error())
		}
	}()

	logrus.Print("Server started at " + port)

	//Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
