package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	tem2024 "github.com/Relax-FM/Test_effective_mobile_go_2024"
	"github.com/Relax-FM/Test_effective_mobile_go_2024/pkg/handler"
	"github.com/Relax-FM/Test_effective_mobile_go_2024/pkg/repository"
	"github.com/Relax-FM/Test_effective_mobile_go_2024/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// @title 				Music library info
// @version 			0.0.1
// @description 		RestAPI Server for Music library

// @contact.name 		Relax FM
// @contact.url 		https://github.com/Relax-FM
// @contact.email 		arton.2@mail.ru

//	@host				localhost:8080
//	@BasePath			/

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Password: 	os.Getenv("DB_PASSWORD"),
		Host: 		os.Getenv("DB_HOST"),
		Port: 		os.Getenv("DB_PORT"),
		Username: 	os.Getenv("DB_USERNAME"),
		DBName: 	os.Getenv("DB_NAME"),
		SSLMode: 	os.Getenv("SSL_MODE"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	rtm, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		logrus.Fatalf("Failed to initialize server: %s", err.Error())
	}

	wtm, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		logrus.Fatalf("Failed to initialize server: %s", err.Error())
	}

	srv := new(tem2024.Server)
	go func() {
		if err:=srv.Run(os.Getenv("APP_PORT"), time.Duration(rtm), time.Duration(wtm), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Music library Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Print("Music library Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil{
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db close: %s", err.Error())
	}
}