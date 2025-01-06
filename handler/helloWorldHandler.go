package handler

import (
	"building-microservice-go/structs"
	"encoding/json"
	"log"
	"net/http"
)

type HelloWorldHandler struct{}

func NewHelloWorldHandler() http.Handler {
	return HelloWorldHandler{}
}

func (h HelloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	type validationContextKey string
	log.Println(r.Context().Value(validationContextKey("name")))
	value := r.Context().Value(validationContextKey("name"))
  var request *structs.HelloWorldRequest
  var ok bool
	if request, ok = value.(*structs.HelloWorldRequest); !ok {
		http.Error(rw, "Invalid request data", http.StatusBadRequest)
		return
	}
	log.Println(request)
	response := structs.HelloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
