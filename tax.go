package taxjar

type Tax struct {
	Breakdown        Breakdown `json:"breakdown"`
	OrderTotalAmount float64   `json:"order_total_amount"`
	Shipping         float64   `json:"shipping"`
	TaxableAmount    float64   `json:"taxable_amount"`
	Rate             float64   `json:"rate"`
	AmountToCollect  float64   `json:"amount_to_collect"`
	HasNexus         bool      `json:"has_nexus"`
	FreightTaxable   bool      `json:"freight_taxable"`
	TaxSource        string    `json:"tax_source"`
}

type TaxList struct {
	Tax Tax `json:"tax"`
}

type Address struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Zip     string `json:"zip,omitempty"`
	Country string `json:"country,omitempty"`
}

type taxParams struct {
	FromCountry    string     `json:"from_country"`
	FromZip        string     `json:"from_zip"`
	FromState      string     `json:"from_state,omitempty"`
	FromCity       string     `json:"from_city,omitempty"`
	FromStreet     string     `json:"from_street,omitempty"`
	ToCountry      string     `json:"to_country,omitempty"`
	ToZip          string     `json:"to_zip"`
	ToState        string     `json:"to_state"`
	ToStreet       string     `json:"to_street,omitempty"`
	ToCity         string     `json:"to_city,omitempty"`
	Shipping       float64    `json:"shipping"`
	Amount         float64    `json:"amount,omitempty"`
	LineItems      []LineItem `json:"line_items,omitempty"`
	NexusAddresses []Address  `json:"nexus_addresses,omitempty"`
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
func (s *TaxService) Calculate(from, to Address, shipping, amount float64) (Tax, error) {
	return s.Repository.get(taxParams{
		FromStreet:  from.Street,
		FromCity:    from.City,
		FromState:   from.State,
		FromZip:     from.Zip,
		FromCountry: from.Country,
		ToStreet:    to.Street,
		ToCity:      to.City,
		ToState:     to.State,
		ToZip:       to.Zip,
		ToCountry:   to.Country,
		Shipping:    shipping,
		Amount:      amount,
	})
}

func (s *TaxService) CalculateItems(from, to Address, nexuses []Address, shipping float64, items []LineItem) (Tax, error) {
	return s.Repository.get(taxParams{
		FromStreet:     from.Street,
		FromCity:       from.City,
		FromState:      from.State,
		FromZip:        from.Zip,
		FromCountry:    from.Country,
		ToStreet:       to.Street,
		ToCity:         to.City,
		ToState:        to.State,
		ToZip:          to.Zip,
		ToCountry:      to.Country,
		Shipping:       shipping,
		LineItems:      items,
		NexusAddresses: nexuses,
	})
}

type TaxLineItem struct {
	Id string `json:"id"`

	// For US transactions
	StateTaxableAmount   float64 `json:"state_taxable_amount"`
	StateSalesTaxRate    float64 `json:"state_sales_tax_rate"`
	StateAmount          float64 `json:"state_amount"`
	CountyTaxableAmount  float64 `json:"county_taxable_amount"`
	CountyTaxRate        float64 `json:"county_tax_rate"`
	CountyAmount         float64 `json:"county_amount"`
	CityTaxableAmount    float64 `json:"city_taxable_amount"`
	CityTaxRate          float64 `json:"city_tax_rate"`
	CityAmount           float64 `json:"city_amount"`
	SpecialTaxableAmount float64 `json:"special_district_taxable_amount"`
	SpecialTaxRate       float64 `json:"special_tax_rate"`
	SpecialAmount        float64 `json:"special_district_amount"`
	TaxCollectable       float64 `json:"tax_collectable"`
	TaxableAmount        float64 `json:"taxable_amount"`

	// For CA transactions
	GstTaxableAmount float64 `json:"gst_taxable_amount"`
	GstTaxRate       float64 `json:"gst_tax_rate"`
	GstAmount        float64 `json:"gst"`
	PstTaxableAmount float64 `json:"pst_taxable_amount"`
	PstTaxRate       float64 `json:"pst_tax_rate"`
	PstAmount        float64 `json:"pst"`
	QstTaxableAmount float64 `json:"qst_taxable_amount"`
	QstTaxRate       float64 `json:"qst_tax_rate"`
	QstAmount        float64 `json:"qst"`
}

type Shipping struct {
	StateTaxableAmount   float64 `json:"state_taxable_amount"`
	StateSalesTaxRate    float64 `json:"state_sales_tax_rate"`
	StateAmount          float64 `json:"state_amount"`
	CountyTaxableAmount  float64 `json:"county_taxable_amount"`
	CountyTaxRate        float64 `json:"county_tax_rate"`
	CountyAmount         float64 `json:"county_amount"`
	CityTaxableAmount    float64 `json:"city_taxable_amount"`
	CityTaxRate          float64 `json:"city_tax_rate"`
	CityAmount           float64 `json:"city_amount"`
	SpecialTaxableAmount float64 `json:"special_district_taxable_amount"`
	SpecialTaxRate       float64 `json:"special_tax_rate"`
	SpecialAmount        float64 `json:"special_district_amount"`

	// For CA transactions
	GstTaxableAmount float64 `json:"gst_taxable_amount"`
	GstTaxRate       float64 `json:"gst_tax_rate"`
	GstAmount        float64 `json:"gst"`
	PstTaxableAmount float64 `json:"pst_taxable_amount"`
	PstTaxRate       float64 `json:"pst_tax_rate"`
	PstAmount        float64 `json:"pst"`
	QstTaxableAmount float64 `json:"qst_taxable_amount"`
	QstTaxRate       float64 `json:"qst_tax_rate"`
	QstAmount        float64 `json:"qst"`
}

type Breakdown struct {
	Shipping  Shipping      `json:"shipping"`
	LineItems []TaxLineItem `json:"line_items"`

	TaxCollectable float64 `json:"tax_collectable"`
	TaxableAmount  float64 `json:"taxable_amount"`

	// For US transactions
	StateTaxableAmount    float64 `json:"state_taxable_amount"`
	StateTaxRate          float64 `json:"state_tax_rate"`
	StateTaxCollectable   float64 `json:"state_tax_collectable"`
	CountyTaxableAmount   float64 `json:"county_taxable_amount"`
	CountyTaxRate         float64 `json:"county_tax_rate"`
	CountyTaxCollectable  float64 `json:"county_tax_collectable"`
	CityTaxableAmount     float64 `json:"city_taxable_amount"`
	CityTaxRate           float64 `json:"city_tax_rate"`
	CityTaxCollectable    float64 `json:"city_tax_collectable"`
	SpecialTaxableAmount  float64 `json:"special_district_taxable_amount"`
	SpecialTaxRate        float64 `json:"special_tax_rate"`
	SpecialTaxCollectable float64 `json:"special_district_tax_collectable"`

	// For CA transactions
	GstTaxableAmount  float64 `json:"gst_taxable_amount"`
	GstTaxCollectable float64 `json:"gst"`
	GstTaxRate        float64 `json:"gst_tax_rate"`
	PstTaxableAmount  float64 `json:"pst_taxable_amount"`
	PstTaxCollectable float64 `json:"pst"`
	PstTaxRate        float64 `json:"pst_tax_rate"`
	QstTaxableAmount  float64 `json:"qst_taxable_amount"`
	QstTaxCollectable float64 `json:"qst"`
	QstTaxRate        float64 `json:"qst_tax_rate"`
}
