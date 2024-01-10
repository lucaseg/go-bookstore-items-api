package services

import (
	"github.com/lucaseg/go-bookstore-items-api/domain/item"
	"github.com/lucaseg/go-bookstore-items-api/domain/queries"
	"github.com/lucaseg/go-bookstore-utils/rest_errors"
)

var (
	ItemService ItemServiceInterface = &itemService{}
)

type ItemServiceInterface interface {
	Create(item.Item) (*item.Item, *rest_errors.RestError)
	GetById(string) (*item.Item, *rest_errors.RestError)
	Search(query queries.EsQuery) ([]item.Item, *rest_errors.RestError)
}

type itemService struct {
}

func (s *itemService) Create(item item.Item) (*item.Item, *rest_errors.RestError) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) GetById(id string) (*item.Item, *rest_errors.RestError) {
	item := item.Item{
		Id: id,
	}

	return &item, nil
}

func (s *itemService) Search(query queries.EsQuery) ([]item.Item, *rest_errors.RestError) {
	item := item.Item{}
	return item.Search(query.Build())
}
