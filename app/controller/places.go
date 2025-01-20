package controller

import (
	"goodmeh/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPlacesController interface {
	Init(*gin.RouterGroup)
}

type PlacesController struct {
	placeService service.IPlaceService
}

func (p *PlacesController) GetPlace(c *gin.Context) {
	// TODO: implement GetPlace
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Not implemented",
	})
}

func (p *PlacesController) Init(r *gin.RouterGroup) {
	g := r.Group("/v1/places")
	g.GET(":id", p.GetPlace)
}

func NewPlacesController(placeService service.IPlaceService) *PlacesController {
	return &PlacesController{}
}
