package item

import (
	"github.com/lucaseg/go-bookstore-items-api/clients/elastic_search"
	"github.com/lucaseg/go-bookstore-utils/rest_errors"
)

const (
	index = "items"
)

func (i *Item) Save() *rest_errors.RestError {
	result, err := elastic_search.Client.Index(index, i)
	if err != nil {
		return rest_errors.InteralServerError("Error trying to save item")
	}
	i.Id = result.Id
	return nil
}
