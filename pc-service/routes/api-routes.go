package routes

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"meogol/pc-service/pc"
	"time"
)

func GetRoutes() *gin.Engine {
	router := gin.New()

	router.ForwardedByClientIP = false
	router.SetTrustedProxies(nil)

	router.Use(ginzap.Ginzap(routesLogger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(routesLogger, true))
	addApiRoutes(router)

	return router
}

func addApiRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	{
		pc.AddRoutes(apiRouter)
	}
}
