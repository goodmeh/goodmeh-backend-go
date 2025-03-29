package router

import (
	"context"
	"goodmeh/app/socket"
	"goodmeh/deps"

	"github.com/gin-gonic/gin"
)

func Init(c *deps.Initialization, ctx context.Context, socketServer *socket.Server) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	socketServer.InitListeners(router)

	api := router.Group("/api")

	c.HealthCtrl.Init(api)
	c.PlacesCtrl.Init(api)

	return router
}
