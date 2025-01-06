package handler

import (
	"building-microservice-go/structs"
	"context"
	"encoding/json"
	"net/http"
)

type ValidationHandler struct {
	next http.Handler
}

func NewValidationHandler(next http.Handler) http.Handler {
	return ValidationHandler{next: next}
}

func (h ValidationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request structs.HelloWorldRequest
	type validationContextKey string
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	if request.Name != "Vigz"{
		http.Error(rw, "Forbidden", http.StatusForbidden)
		return
	}
	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)
	h.next.ServeHTTP(rw, r)
}
