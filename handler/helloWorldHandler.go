package handler

import (
	"building-microservice-go/constants"
	"building-microservice-go/structs"
	"encoding/json"
	"net/http"
)

type HelloWorldHandler struct{}

func NewHelloWorldHandler() http.Handler {
	return HelloWorldHandler{}
}

func (h HelloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	value := r.Context().Value(constants.RequestKey)
	var request *structs.HelloWorldRequest
	var ok bool
	if request, ok = value.(*structs.HelloWorldRequest); !ok {
		http.Error(rw, "Invalid request data", http.StatusBadRequest)
		return
	}
	response := structs.HelloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
