package deps

import (
	"goodmeh/app/controller"
	"goodmeh/app/socket"
)

type Initialization struct {
	HealthCtrl   controller.IHealthController
	PlacesCtrl   controller.IPlacesController
	SocketServer *socket.Server
}

func NewInitialization(
	healthCtrl controller.IHealthController,
	placesCtrl controller.IPlacesController,
	socketServer *socket.Server,
) *Initialization {
	return &Initialization{
		HealthCtrl:   healthCtrl,
		PlacesCtrl:   placesCtrl,
		SocketServer: socketServer,
	}
}
