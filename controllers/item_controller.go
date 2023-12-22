package controllers

import (
	"github.com/lucaseg/go-bookstore-oauth/oauth"
	"net/http"
)

var (
	ItemController itemControllerInterface = &itemController{}
)

type itemControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemController struct {
}

func (c *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return error to caller
		return
	}

}

func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {

}
