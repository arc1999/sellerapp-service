package main

import (
	"api-scraping-service/server"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	config := AppConfig{}

	log.Infoln("starting application")
	config.LoadEnv()
	server.Init()
}