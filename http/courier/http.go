package courier

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/vageeshabr/courier-demo/services"
)

type server struct {
	svc services.CourierService
}

func New(svc services.CourierService) *server {
	return &server{svc: svc}
}

func (s *server) Create(r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	m := struct {
		Id             int
		ItemCount      int
		PackageDetails struct {
			SmallBags int
			LargeBags int
			Weight    float64
		}
		Address struct {
			L1      string
			L2      string
			State   string
			ZipCode int
			Country string
		}
	}{}
	err := json.Unmarshal(b, m)
	if err != nil {
		return nil, errors.New("invalid json body")
	}

	err = s.svc.CreateCourier(m)
	return nil, err
}

func (s *server) Delete() {

}
