package main

import (
	"fmt"
	"meogol/pc-service/config"
	"meogol/pc-service/database"
	"meogol/pc-service/database/pc"
	"meogol/pc-service/logger"
	"meogol/pc-service/routes"
)

func main() {
	srvConf := config.CurrentConfig.Server
	r := routes.GetRoutes()
	logger.Logger.Info("Start server")
	database.InitDatabase()
	pc.CreateTable()
	r.Run(fmt.Sprintf("%s:%d", srvConf.Host, srvConf.Port))
}
