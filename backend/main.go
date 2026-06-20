package main

import (
	"go-trading/backend/app/controllers"
	"go-trading/backend/config"
	"go-trading/backend/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
