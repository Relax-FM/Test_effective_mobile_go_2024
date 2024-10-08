package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	tem2024 "github.com/Relax-FM/Test_effective_mobile_go_2024"
	"github.com/Relax-FM/Test_effective_mobile_go_2024/internal/handler"
	"github.com/Relax-FM/Test_effective_mobile_go_2024/internal/repository"
	"github.com/Relax-FM/Test_effective_mobile_go_2024/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title 				Music library info
// @version 			0.0.1
// @description 		RestAPI Server for Music library

// @contact.name 		Relax FM
// @contact.url 		https://github.com/Relax-FM
// @contact.email 		arton.2@mail.ru

//	@host				localhost:8800
//	@BasePath			/

func main() {	
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	ll, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
    if err != nil {
        ll = logrus.DebugLevel
    }
    logrus.SetLevel(ll)

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

	defer logrus.Print("Music library Shuted Down")
	defer srv.Shutdown(context.Background())
	defer repository.CloseDB(db)
	defer logrus.Print("Music library Shutting Down")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<- quit
 
}