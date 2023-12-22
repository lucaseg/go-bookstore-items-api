package services

import (
	"github.com/lucaseg/go-bookstore-items-api/domain/item"
	"github.com/lucaseg/go-bookstore-utils/rest_errors"
)

var (
	ItemService ItemServiceInterface = &itemService{}
)

type ItemServiceInterface interface {
	Create(item.Item) (*item.Item, *rest_errors.RestError)
	GetById(string) (*item.Item, *rest_errors.RestError)
}

type itemService struct {
}

func (s *itemService) Create(item item.Item) (*item.Item, *rest_errors.RestError) {
	return nil, nil
}

func (s *itemService) GetById(id string) (*item.Item, *rest_errors.RestError) {
	return nil, nil
}
