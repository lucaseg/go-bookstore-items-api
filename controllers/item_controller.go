package controllers

import (
	"encoding/json"
	"github.com/lucaseg/go-bookstore-items-api/domain/item"
	"github.com/lucaseg/go-bookstore-items-api/services"
	"github.com/lucaseg/go-bookstore-items-api/utils/http_utils"
	"github.com/lucaseg/go-bookstore-oauth/oauth"
	"github.com/lucaseg/go-bookstore-utils/rest_errors"
	"io/ioutil"
	"net/http"
	"strconv"
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
		http_utils.ResponseError(w, (*rest_errors.RestError)(err))
		return
	}

	var newItem item.Item
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restError := rest_errors.BadRequest("invalid json body")
		http_utils.ResponseError(w, restError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &newItem); err != nil {
		restError := rest_errors.BadRequest("invalid json body")
		http_utils.ResponseError(w, restError)
		return
	}

	newItem.Seller = strconv.FormatInt(oauth.GetCallerId(r), 10)

	result, errSrv := services.ItemService.Create(newItem)
	if errSrv != nil {
		http_utils.ResponseError(w, errSrv)
		return
	}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}

func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {

}
