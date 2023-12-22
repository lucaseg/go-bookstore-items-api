package application

import (
	"github.com/lucaseg/go-bookstore-items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.PingController.Get).Methods(http.MethodGet)

}
