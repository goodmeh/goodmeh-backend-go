// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package deps

import (
	"context"
	"github.com/google/wire"
	"goodmeh/app/controller"
	"goodmeh/app/events"
	"goodmeh/app/repository"
	"goodmeh/app/service"
	"goodmeh/app/socket"
)

// Injectors from injector.go:

func Initialize(db repository.DBTX, ctx context.Context, socketServer *socket.Server) *Initialization {
	queries := ProvideQueries(db)
	eventBus := events.NewEventBus()
	fieldService := service.NewFieldService(ctx, queries, eventBus)
	healthController := controller.NewHealthController(fieldService)
	placeService := service.NewPlaceService(ctx, queries, eventBus)
	reviewService := service.NewReviewService(ctx, queries, eventBus)
	placesController := controller.NewPlacesController(placeService, reviewService, socketServer, eventBus)
	initialization := NewInitialization(healthController, placesController, socketServer)
	return initialization
}

// injector.go:

func ProvideQueries(db repository.DBTX) *repository.Queries {
	return repository.New(db)
}

var eventBusSet = wire.NewSet(events.NewEventBus)

var repositorySet = wire.NewSet(
	ProvideQueries,
)

var placeServiceSet = wire.NewSet(service.NewPlaceService, wire.Bind(new(service.IPlaceService), new(*service.PlaceService)))

var reviewServiceSet = wire.NewSet(service.NewReviewService, wire.Bind(new(service.IReviewService), new(*service.ReviewService)))

var fieldServiceSet = wire.NewSet(service.NewFieldService, wire.Bind(new(service.IFieldService), new(*service.FieldService)))

var healthControllerSet = wire.NewSet(controller.NewHealthController, wire.Bind(new(controller.IHealthController), new(*controller.HealthController)))

var placesControllerSet = wire.NewSet(controller.NewPlacesController, wire.Bind(new(controller.IPlacesController), new(*controller.PlacesController)))
