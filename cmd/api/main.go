package main

import (
	"context"
	"github.com/AssylzhanZharzhanov/dochq-test-task/db/postgres"
	grpcHandler "github.com/AssylzhanZharzhanov/dochq-test-task/pkg/handler/grpc"
	restHandler "github.com/AssylzhanZharzhanov/dochq-test-task/pkg/handler/rest"
	grpcRepository "github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository/grpc"
	restRepository "github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository/rest"
	grpcService "github.com/AssylzhanZharzhanov/dochq-test-task/pkg/service/grpc"
	restService "github.com/AssylzhanZharzhanov/dochq-test-task/pkg/service/rest"
	grpcServer "github.com/AssylzhanZharzhanov/dochq-test-task/server/grpc"
	"github.com/AssylzhanZharzhanov/dochq-test-task/server/rest"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	restSrv := new(rest.Server)
	grpcSrv := grpcServer.NewServer()

	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error in reading env file: %s", err.Error())
	}

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

	restRepositories := restRepository.NewRepository(db)
	restServices := restService.NewService(restRepositories)
	restHandlers := restHandler.NewHandler(restServices)

	restPort := viper.GetString("server.port")
	go func() {
		if err := restSrv.Run(restPort, restHandlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error in starting rest server: %s", err.Error())
		}
	}()
	logrus.Print("rest server started at " + restPort)

	grpcRepositories := grpcRepository.NewRepository(db)
	grpcServices := grpcService.NewService(grpcRepositories)
	grpcHandlers := grpcHandler.NewHandler(grpcServices)

	grpcPort := viper.GetString("grpc.port")
	go func() {
		if err := grpcSrv.Run(grpcPort, grpcHandlers); err != nil {
			logrus.Fatalf("Error in starting grpc server: %s", err.Error())
		}
	}()

	logrus.Print("grpc server started at " + grpcPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := restSrv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	grpcSrv.Stop()

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
