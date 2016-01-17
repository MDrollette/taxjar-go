package taxjar

import (
	"encoding/json"
)

// CategoryRepository defines the interface for working with Categories through the API.
type CategoryRepository interface {
	list(categoryListParams) (CategoryList, error)
}

// CategoryApi implements CategoryRepository
type CategoryApi struct {
	client *Client
}

func (api CategoryApi) list(params categoryListParams) (CategoryList, error) {
	categoryList := CategoryList{}
	data, err := api.client.Get("/categories", params)
	if err != nil {
		return categoryList, err
	}
	err = json.Unmarshal(data, &categoryList)
	return categoryList, err
}
