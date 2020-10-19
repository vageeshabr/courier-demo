package courier

import (
	"github.com/vageeshabr/courier-demo/models"
	"github.com/vageeshabr/courier-demo/services"
)

type service struct {
	domesticCSP      services.CourierServiceProvider
	internationalCSP services.CourierServiceProvider
}

func New(d, i services.CourierServiceProvider) *service {
	return &service{domesticCSP: d, internationalCSP: i}
}

func (s *service) CreateCourier(o *models.Order) error {
	domestic := s.isDomestic(o)
	var p services.CourierServiceProvider
	if domestic {
		p = s.domesticCSP
	} else {
		p = s.internationalCSP
	}
	return p.SendCourier(o)
}

func (s *service) isDomestic(o *models.Order) bool {
	return o.Address.Country == "IN"
}
