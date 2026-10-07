package main

import (
	"log"
	"net/http"

	"order/routers"
)

func main() {
	r := routers.New()

	log.Println("cerver run")
	log.Fatal(http.ListenAndServe(":8080", r))
}
