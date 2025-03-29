package router

import (
	"context"
	"goodmeh/deps"

	"github.com/gin-gonic/gin"
)

func Init(c *deps.Initialization, ctx context.Context) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	c.SocketServer.InitListeners(router)

	api := router.Group("/api")

	c.HealthCtrl.Init(api)
	c.PlacesCtrl.Init(api)

	return router
}
