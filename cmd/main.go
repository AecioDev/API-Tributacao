package main

import (
	"api-tributacao/config"
	"api-tributacao/src/controllers/controller"
	"api-tributacao/src/db"
	"api-tributacao/src/globals"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalln("Failed to parse configuration: ", err)
	}

	database := db.New(cfg)
	defer db.CloseConnection(database)

	globals.SetDev(true)

	sys := controller.New(gin.Default(), database, cfg)

	sys.StartServer(cfg.Server.Port)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit
}
