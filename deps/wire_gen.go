// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package deps

import (
	"context"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5"
	"goodmeh/app/controller"
	"goodmeh/app/repository"
	"goodmeh/app/service"
)

// Injectors from injector.go:

func Initialize(db *pgx.Conn, ctx context.Context) *Initialization {
	healthController := controller.NewHealthController()
	queries := ProvideQueries(db)
	placeService := service.NewPlaceService(ctx, queries)
	placesController := controller.NewPlacesController(placeService)
	initialization := NewInitialization(healthController, placesController)
	return initialization
}

// injector.go:

func ProvideQueries(db *pgx.Conn) *repository.Queries {
	return repository.New(db)
}

var repositorySet = wire.NewSet(
	ProvideQueries,
)

var placeServiceSet = wire.NewSet(service.NewPlaceService, wire.Bind(new(service.IPlaceService), new(*service.PlaceService)))

var healthControllerSet = wire.NewSet(controller.NewHealthController, wire.Bind(new(controller.IHealthController), new(*controller.HealthController)))

var placesControllerSet = wire.NewSet(controller.NewPlacesController, wire.Bind(new(controller.IPlacesController), new(*controller.PlacesController)))
