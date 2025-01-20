//go:build wireinject
// +build wireinject

package config

import (
	"goodmeh/app/controller"
	"goodmeh/app/service"

	"github.com/google/wire"
)

var placeServiceSet = wire.NewSet(service.NewPlaceService,
	wire.Bind(new(service.IPlaceService), new(*service.PlaceService)))

var healthControllerSet = wire.NewSet(controller.NewHealthController,
	wire.Bind(new(controller.IHealthController), new(*controller.HealthController)))

var placesControllerSet = wire.NewSet(controller.NewPlacesController,
	wire.Bind(new(controller.IPlacesController), new(*controller.PlacesController)))

func Initialize() *Initialization {
	wire.Build(NewInitialization, healthControllerSet, placesControllerSet, placeServiceSet)
	return nil
}
