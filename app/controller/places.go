package controller

import (
	"goodmeh/app/dto/response"
	"goodmeh/app/events"
	"goodmeh/app/mapper"
	"goodmeh/app/service"
	"goodmeh/app/socket"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goodmeh/backend-private/models"
)

type IPlacesController interface {
	Init(*gin.RouterGroup)
}

type PlacesController struct {
	placeService  service.IPlaceService
	reviewService service.IReviewService
	socketServer  *socket.Server
	eventBus      *events.EventBus
}

func (p *PlacesController) RequestPlace(c *gin.Context) {
	id := c.Param("id")
	place, err := p.placeService.RequestPlace(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if place == nil {
		c.JSON(http.StatusOK, response.RequestPlaceResponseDto{
			Status: "Scraping",
		})
		return
	}
	c.JSON(http.StatusOK, mapper.ToPlaceResponseDto(*place))
}

func (p *PlacesController) GetRandomPlaces(c *gin.Context) {
	placesModel, err := p.placeService.GetRandomPlaces()
	places := mapper.ToPlacePreviewResponseDtos(placesModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, places)
}

func (p *PlacesController) GetPlaceReviews(c *gin.Context) {
	id := c.Param("id")
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid page",
		})
		return
	}
	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid per_page",
		})
		return
	}
	reviewsModel, err := p.placeService.GetPlaceReviews(id, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	reviewIds := make([]string, len(reviewsModel))
	for i, review := range reviewsModel {
		reviewIds[i] = review.Review.ID
	}
	imageUrls, err := p.reviewService.GetReviewsImages(reviewIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.GetPlaceReviewsResponseDto{
		Data:    mapper.ToReviewResponseDto(reviewsModel, imageUrls, perPage),
		HasNext: len(reviewsModel) == perPage,
	})
}

func (p *PlacesController) GetPlaceImages(c *gin.Context) {
	id := c.Param("id")
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid page",
		})
		return
	}
	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid per_page",
		})
		return
	}
	imageUrls, err := p.placeService.GetPlaceImages(id, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.GetPlaceImagesResponseDto{
		Data:    imageUrls,
		HasNext: len(imageUrls) == perPage,
	})
}

func (p *PlacesController) GetPlaceNames(c *gin.Context) {
	placeNames, err := p.placeService.GetPlaceNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, placeNames)
}

func (p *PlacesController) OnPlaceScrape(place models.Place) error {
	return p.socketServer.To(place.ID, place)
}

func (p *PlacesController) Init(r *gin.RouterGroup) {
	g := r.Group("/v1/places")
	g.GET("/", p.GetPlaceNames)
	g.POST("/:id", p.RequestPlace)
	g.GET("/:id/reviews", p.GetPlaceReviews)
	g.GET("/:id/images", p.GetPlaceImages)
	g.GET("/discover", p.GetRandomPlaces)

	p.eventBus.Subscribe(events.ON_PLACE_SCRAPE, events.AssertHandler(p.OnPlaceScrape))
}

func NewPlacesController(
	placeService service.IPlaceService,
	reviewService service.IReviewService,
	socketServer *socket.Server,
	eventBus *events.EventBus,
) *PlacesController {
	return &PlacesController{
		placeService,
		reviewService,
		socketServer,
		eventBus,
	}
}
