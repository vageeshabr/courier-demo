package bluedart

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/vageeshabr/courier-demo/models"
)

type bluedart struct {
	APIURL string
	User   string
	Pass   string
}

func (b *bluedart) SendCourier(order *models.Order) error {
	o, _ := json.Marshal(order)
	bb := bytes.NewBuffer(o)
	_, err := http.Post(b.APIURL, "application/json", bb)
	return err
}

func (b *bluedart) CancelCourier(order *models.Order) error {
	panic("implement me")
}

func (b *bluedart) GetCurrentLocation(order *models.Order) (*models.Location, error) {
	panic("implement me")
}

func New(apiURL, user, pass string) *bluedart {
	return &bluedart{
		APIURL: "",
		User:   "",
		Pass:   "",
	}
}
