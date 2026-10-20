package main

import (
	"fmt"
	"net/http"
	"order/routers"
)

func main() {
	fmt.Println("server run")
	http.ListenAndServe(":8081", routers.New())
}
