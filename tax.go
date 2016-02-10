package taxjar

type Tax struct {
	Breakdown        Breakdown `json:"breakdown"`
	OrderTotalAmount float64   `json:"order_total_amount"`
	Shipping         float64   `json:"shipping"`
	TaxableAmount    float64   `json:"taxable_amount"`
	AmountToCollect  float64   `json:"amount_to_collect"`
	HasNexus         bool      `json:"has_nexus"`
	FreightTaxable   bool      `json:"freight_taxable"`
	TaxSource        string    `json:"tax_source"`
}

type TaxList struct {
	Tax Tax `json:"tax"`
}

type taxParams struct {
	FromCountry string     `json:"from_country"`
	FromZip     string     `json:"from_zip"`
	FromState   string     `json:"from_state,omitempty"`
	FromCity    string     `json:"from_city,omitempty"`
	FromStreet  string     `json:"from_street,omitempty"`
	ToCountry   string     `json:"to_country,omitempty"`
	ToZip       string     `json:"to_zip"`
	ToState     string     `json:"to_state"`
	Shipping    float64    `json:"shipping"`
	Amount      float64    `json:"amount,omitempty"`
	LineItems   []LineItem `json:"line_items,omitempty"`
}

type LineItem struct {
	Id             string  `json:"id,omitempty"`
	Quantity       int64   `json:"quantity,omitempty"`
	ProductTaxCode string  `json:"product_tax_code,omitempty"`
	UnitPrice      float64 `json:"unit_price,omitempty"`
	Discount       float64 `json:"discount,omitempty"`
}

type TaxService struct {
	Repository TaxRepository
}

// Calculate sales Tax for a given order
func (s *TaxService) Calculate(fromStreet, fromCity, fromState, fromZip, fromCountry, toState, toZip, toCountry string, shipping, amount float64) (Tax, error) {
	return s.Repository.get(taxParams{
		FromStreet:  fromStreet,
		FromCity:    fromCity,
		FromState:   fromState,
		FromZip:     fromZip,
		FromCountry: fromCountry,
		ToState:     toState,
		ToZip:       toZip,
		ToCountry:   toCountry,
		Shipping:    shipping,
		Amount:      amount,
	})
}

func (s *TaxService) CalculateItems(fromStreet, fromCity, fromState, fromZip, fromCountry, toState, toZip, toCountry string, shipping float64, items []LineItem) (Tax, error) {
	return s.Repository.get(taxParams{
		FromStreet:  fromStreet,
		FromCity:    fromCity,
		FromState:   fromState,
		FromZip:     fromZip,
		FromCountry: fromCountry,
		ToState:     toState,
		ToZip:       toZip,
		ToCountry:   toCountry,
		Shipping:    shipping,
		LineItems:   items,
	})
}

type TaxLineItem struct {
	Id                           string  `json:"id"`
	StateTaxableAmount           float64 `json:"state_taxable_amount"`
	StateSalesTaxRate            float64 `json:"state_sales_tax_rate"`
	CountyTaxableAmount          float64 `json:"county_taxable_amount"`
	CountyTaxRate                float64 `json:"county_tax_rate"`
	CityTaxableAmount            float64 `json:"city_taxable_amount"`
	CityTaxRate                  float64 `json:"city_tax_rate"`
	SpecialDistrictTaxableAmount float64 `json:"special_district_taxable_amount"`
	SpecialTaxRate               float64 `json:"special_tax_rate"`
}

type Shipping struct {
	StateAmount           float64 `json:"state_amount"`
	StateSalesTaxRate     float64 `json:"state_sales_tax_rate"`
	CountyAmount          float64 `json:"county_amount"`
	CountyTaxRate         float64 `json:"county_tax_rate"`
	CityAmount            float64 `json:"city_amount"`
	CityTaxRate           float64 `json:"city_tax_rate"`
	SpecialDistrictAmount float64 `json:"special_district_amount"`
	SpecialTaxRate        float64 `json:"special_tax_rate"`
}

type Breakdown struct {
	Shipping                      Shipping      `json:"shipping"`
	LineItems                     []TaxLineItem `json:"line_items"`
	StateTaxableAmount            float64       `json:"state_taxable_amount"`
	StateTaxCollectable           float64       `json:"state_tax_collectable"`
	CountyTaxableAmount           float64       `json:"county_taxable_amount"`
	CountyTaxCollectable          float64       `json:"county_tax_collectable"`
	CityTaxableAmount             float64       `json:"city_taxable_amount"`
	CityTaxCollectable            float64       `json:"city_tax_collectable"`
	SpecialDistrictTaxableAmount  float64       `json:"special_district_taxable_amount"`
	SpecialDistrictTaxCollectable float64       `json:"special_district_tax_collectable"`
}
