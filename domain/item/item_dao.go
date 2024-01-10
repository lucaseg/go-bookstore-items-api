package item

import (
	"encoding/json"
	"fmt"
	"github.com/lucaseg/go-bookstore-items-api/clients/elastic_search"
	"github.com/lucaseg/go-bookstore-utils/rest_errors"
	"github.com/olivere/elastic"
	"strings"
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

func (i *Item) Get() *rest_errors.RestError {
	result, err := elastic_search.Client.Get(index, "_doc", i.Id)
	if err != nil {
		if strings.Contains("404", err.Error()) {
			return rest_errors.NotFound(fmt.Sprintf("Document id : %s, not found", i.Id))
		}
		return rest_errors.InteralServerError(fmt.Sprintf("Error trying to get document with id : %s", i.Id))
	}
	bytes, err := result.Source.MarshalJSON()
	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.InteralServerError("error trying to deserialize document")
	}
	return nil
}

func (i *Item) Search(query elastic.Query) ([]Item, *rest_errors.RestError) {
	result, err := elastic_search.Client.Search(index, query)
	if err != nil {
		return nil, rest_errors.InteralServerError("Error trying to search documents")
	}

	items := make([]Item, result.TotalHits())
	for index, value := range result.Hits.Hits {
		bytes, _ := value.Source.MarshalJSON()

		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.InteralServerError("")
		}
		item.Id = value.Id
		items[index] = item
	}
	return items, nil
}
