package taxjar

// Rate defines the returned object for rate requests
type Rate struct {
	Zip                   string  `json:"zip"`
	State                 string  `json:"state"`
	StateRate             float64 `json:"state_rate,string"`
	County                string  `json:"county"`
	CountyRate            float64 `json:"county_rate,string"`
	City                  string  `json:"city"`
	CityRate              float64 `json:"city_rate,string"`
	CombinedDistrictRate  float64 `json:"combined_district_rate,string"`
	CombinedRate          float64 `json:"combined_rate,string"`
	Country               string  `json:"country"`
	Name                  string  `json:"name"`
	StandardRate          float64 `json:"standard_rate,string"`
	ReducedRate           float64 `json:"reduced_rate,string"`
	SuperReducedRate      float64 `json:"super_reduced_rate,string"`
	ParkingRate           float64 `json:"parking_rate,string"`
	DistanceSaleThreshold float64 `json:"distance_sale_threshold,string"`
	FreightTaxable        *bool   `json:"freight_taxable"`
}

// RateList is the wrapper for the actual rate object
type RateList struct {
	Rate Rate `json:"rate"`
}

type rateParams struct {
	Country string `url:"country,omitempty"`
	Zip     string `url:"-"`
	State   string `url:"state,omitempty"`
	City    string `url:"city,omitempty"`
	Street  string `url:"street,omitempty"`
}

// RateCountry returns an option which sets the `country` parameter
func RateCountry(country string) func(*rateParams) error {
	return func(rp *rateParams) error {
		rp.Country = country
		return nil
	}
}

// RateState returns an option which sets the `state` parameter
func RateState(state string) func(*rateParams) error {
	return func(rp *rateParams) error {
		rp.State = state
		return nil
	}
}

// RateCity returns an option which sets the `city` parameter
func RateCity(city string) func(*rateParams) error {
	return func(rp *rateParams) error {
		rp.City = city
		return nil
	}
}

// RateStreet returns an option which sets the `street` parameter
func RateStreet(street string) func(*rateParams) error {
	return func(rp *rateParams) error {
		rp.Street = street
		return nil
	}
}

// RateService interfaces with the rates part of the API
type RateService struct {
	Repository RateRepository
}

// Get a Rate
func (s *RateService) Get(zip string, options ...func(*rateParams) error) (Rate, error) {
	params := rateParams{Zip: zip}
	for _, option := range options {
		if err := option(&params); nil != err {
			return Rate{}, err
		}
	}

	return s.Repository.get(params)
}
