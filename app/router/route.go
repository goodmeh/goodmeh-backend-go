package router

import (
	"goodmeh/config"

	"github.com/gin-gonic/gin"
)

func Init(c *config.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")

	c.HealthCtrl.Init(api)
	c.PlacesCtrl.Init(api)

	return router
}
