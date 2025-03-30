//go:build wireinject
// +build wireinject

package deps

import (
	"context"
	"goodmeh/app/controller"
	"goodmeh/app/events"
	"goodmeh/app/repository"
	"goodmeh/app/service"
	"goodmeh/app/socket"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5"
)

func ProvideQueries(db *pgx.Conn) *repository.Queries {
	return repository.New(db)
}

var eventBusSet = wire.NewSet(
	events.NewEventBus,
)

var repositorySet = wire.NewSet(
	ProvideQueries,
)

var placeServiceSet = wire.NewSet(service.NewPlaceService,
	wire.Bind(new(service.IPlaceService), new(*service.PlaceService)))

var reviewServiceSet = wire.NewSet(service.NewReviewService,
	wire.Bind(new(service.IReviewService), new(*service.ReviewService)))

var fieldServiceSet = wire.NewSet(service.NewFieldService,
	wire.Bind(new(service.IFieldService), new(*service.FieldService)))

var healthControllerSet = wire.NewSet(controller.NewHealthController,
	wire.Bind(new(controller.IHealthController), new(*controller.HealthController)))

var placesControllerSet = wire.NewSet(controller.NewPlacesController,
	wire.Bind(new(controller.IPlacesController), new(*controller.PlacesController)))

func Initialize(db *pgx.Conn, ctx context.Context, socketServer *socket.Server) *Initialization {
	wire.Build(
		NewInitialization,
		healthControllerSet,
		placesControllerSet,
		placeServiceSet,
		reviewServiceSet,
		fieldServiceSet,
		repositorySet,
		eventBusSet,
	)
	return nil
}
