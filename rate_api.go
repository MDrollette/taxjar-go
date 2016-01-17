package taxjar

import (
	"encoding/json"
)

// RateRepository defines the interface for working with Rates through the API.
type RateRepository interface {
	get(rateParams) (Rate, error)
}

// RateApi implements RateRepository
type RateApi struct {
	client *Client
}

func (api RateApi) get(params rateParams) (Rate, error) {
	rateList := RateList{}
	data, err := api.client.Get("/rates/"+params.Zip, params)
	if err != nil {
		return rateList.Rate, err
	}
	err = json.Unmarshal(data, &rateList)
	return rateList.Rate, err
}
