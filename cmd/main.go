package main

import (
	"testbank/config"
	"testbank/pkg/handler"
	"testbank/pkg/repository"
	"testbank/pkg/service"
	"testbank/router"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(".env file dont load :", err)
	}

	c := config.GetConfig()

	db, err := repository.PostgresDBConnect(c)
	if err != nil {
		logrus.Fatal(err)
	}

	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	router := router.NewRouter(h)
	router.Start()

}
