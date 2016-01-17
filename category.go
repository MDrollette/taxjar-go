package taxjar

type Category struct {
	Name           string `json:"name"`
	ProductTaxCode string `json:"product_tax_code"`
	Description    string `json:"description"`
}

type CategoryList struct {
	Categories []Category `json:"categories"`
}

type categoryListParams struct{}

type CategoryService struct {
	Repository CategoryRepository
}

// List all Categories
func (s *CategoryService) List() (CategoryList, error) {
	return s.Repository.list(categoryListParams{})
}
