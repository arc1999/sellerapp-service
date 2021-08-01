package main

import (
	"api-productinfo-service/db"
	"api-productinfo-service/server"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	config := AppConfig{}

	log.Infoln("starting application")
	config.LoadEnv()
	db.InitDb()
	server.Init()

}
