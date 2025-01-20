package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHealthController interface {
	Health(c *gin.Context)

	Init(*gin.RouterGroup)
}

type HealthController struct{}

func (h *HealthController) Health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func (h *HealthController) Init(r *gin.RouterGroup) {
	g := r.Group("/v1/health")
	g.GET("", h.Health)
}

func NewHealthController() *HealthController {
	return &HealthController{}
}
