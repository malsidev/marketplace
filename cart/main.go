package main

import (
	"log"
	"net/http"

	"order/config"
	"order/database"
	"order/models"
	"order/routers"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	_ = db.Migrator().DropTable(&models.Carts{}, &models.CartItem{})
	if err := db.AutoMigrate(&models.Carts{}, &models.CartItem{}); err != nil {
		log.Fatal(err)
	}

	log.Println("database connected", db != nil)
	r := routers.New(db)

	log.Println("server run")
	log.Fatal(http.ListenAndServe(":8080", r))
}
