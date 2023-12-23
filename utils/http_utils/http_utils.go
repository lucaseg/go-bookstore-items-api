package http_utils

import (
	"encoding/json"
	"github.com/lucaseg/go-bookstore-utils/rest_errors"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(body)
	w.WriteHeader(statusCode)
}

func ResponseError(w http.ResponseWriter, err *rest_errors.RestError) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(err)
	w.WriteHeader(err.Status)
}
