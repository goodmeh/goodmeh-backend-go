package config

import "goodmeh/app/controller"

type Initialization struct {
	HealthCtrl controller.IHealthController
	PlacesCtrl controller.IPlacesController
}

func NewInitialization(healthCtrl controller.IHealthController, placesCtrl controller.IPlacesController) *Initialization {
	return &Initialization{
		HealthCtrl: healthCtrl,
		PlacesCtrl: placesCtrl,
	}
}
