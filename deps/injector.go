//go:build wireinject
// +build wireinject

package deps

import (
	"context"
	"goodmeh/app/controller"
	"goodmeh/app/repository"
	"goodmeh/app/service"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5"
)

func ProvideQueries(db *pgx.Conn) *repository.Queries {
	return repository.New(db)
}

var repositorySet = wire.NewSet(
	ProvideQueries,
)

var placeServiceSet = wire.NewSet(service.NewPlaceService,
	wire.Bind(new(service.IPlaceService), new(*service.PlaceService)))

var healthControllerSet = wire.NewSet(controller.NewHealthController,
	wire.Bind(new(controller.IHealthController), new(*controller.HealthController)))

var placesControllerSet = wire.NewSet(controller.NewPlacesController,
	wire.Bind(new(controller.IPlacesController), new(*controller.PlacesController)))

func Initialize(db *pgx.Conn, ctx context.Context) *Initialization {
	wire.Build(NewInitialization, healthControllerSet, placesControllerSet, placeServiceSet, repositorySet)
	return nil
}
