package main

import (
	"building-microservice-go/structs"
	"building-microservice-go/handler"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)
	http.Handle("/hello", handler.NewValidationHandler(handler.NewHelloWorldHandler()))
	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	/*
		data, err := json.Marshal(response)
		if err != nil {
			panic("Ooops")
		}
		fmt.Fprint(w, string(data))
		fmt.Fprint(w, response)
	*/
	var request structs.HelloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := structs.HelloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
