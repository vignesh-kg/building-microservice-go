package handler

import (
	"building-microservice-go/constants"
	"building-microservice-go/structs"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type ValidationHandler struct {
	next http.Handler
}

func NewValidationHandler(next http.Handler) http.Handler {
	return ValidationHandler{next: next}
}

func (h ValidationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request structs.HelloWorldRequest
	
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	if !strings.EqualFold(request.Name, "VIGZ"){
		http.Error(rw, "Forbidden", http.StatusForbidden)
		return
	}
	c := context.WithValue(r.Context(), constants.RequestKey, &request)
	r = r.WithContext(c)
	h.next.ServeHTTP(rw, r)
}
