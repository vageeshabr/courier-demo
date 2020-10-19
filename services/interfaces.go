package services

import "github.com/vageeshabr/courier-demo/models"

type CourierServiceProvider interface {
	SendCourier(order *models.Order) error
	CancelCourier(order *models.Order) error
	GetCurrentLocation(order *models.Order) (*models.Location, error)
}

type CourierService interface {
	CreateCourier(order *models.Order) error
	CancelCourier(order *models.Order) error
	Track(order *models.Order) error
}

// our_id, service_provide_id, provider
