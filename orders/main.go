package main

import (
	"fmt"
	"log"
	"net/http"
	"order/config"
	"order/db"
	"order/routers"
)

func main() {
	cfg := config.Load()
	db, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connected", db != nil)
	fmt.Println("server run")
	http.ListenAndServe(":8081", routers.New(db))
}
