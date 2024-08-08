package main

import (
	"go-todo-app/router"
	"log"
	"net/http"
)

func main() {
	router := router.InitializeRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
