package controller

import (
	"goodmeh/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHealthController interface {
	Health(c *gin.Context)

	Init(*gin.RouterGroup)
}

type HealthController struct {
	// This is here because wire complains because it thinks fieldService is not used
	// but it is used to listen to events
	fieldService service.IFieldService
}

func (h *HealthController) Health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func (h *HealthController) Init(r *gin.RouterGroup) {
	g := r.Group("/v1/health")
	g.GET("", h.Health)
}

func NewHealthController(fieldService service.IFieldService) *HealthController {
	return &HealthController{fieldService: fieldService}
}
