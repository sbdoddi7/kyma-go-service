package main

import (
	"log"
	"net/http"

	"github.com/sbdoddi7/kyma-go-service/internal/handler"
)

func main() {
	h := handler.NewHandler()

	http.HandleFunc("/", h.Health)
	http.HandleFunc("/hello", h.Hello)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
