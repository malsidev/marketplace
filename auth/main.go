package main

import (
	"log"
	"net/http"
	"fmt"

	"auth/internal/config"
	"auth/internal/database"
	"auth/internal/middleware"
	"github.com/go-chi/chi/v5"
	"auth/internal/handler"
	"auth/internal/service"
	"auth/internal/repository"
)

func main() {
	
	cfg := config.Load()

	
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database connected:", db != nil)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	jwtService := service.NewJWTService(cfg.JWTSecret)
	UserRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(UserRepository, jwtService)
	AuthHandler := handler.NewAuthHandler(authService)

	// Auth routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/reg", AuthHandler.Register)
		r.Post("/log", AuthHandler.Login)
	})

	log.Println("server started on :8000")

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err)
	}
}
