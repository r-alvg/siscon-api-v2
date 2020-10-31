package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"siscon/configs"
	"siscon/db"
	"siscon/routes"
)

func main() {
	configs.Read()
	db.Start()

	router := routes.Router()
	logrus.Infof("server runing ir port %v path: %s ", configs.Properties.Server.Port, "/siscon/v2")
	logrus.Fatal(http.ListenAndServe(configs.Properties.Server.Port, router))
}
