package main

import (
	"log"
	"net/http"

	"auth/internal/config"
	"auth/internal/database"
	"auth/internal/middleware"
	"github.com/go-chi/chi/v5"
	"auth/internal/handler"
	"auth/internal/service"
	"auth/internal/repository"
)

func main() {
	// Загружаем конфигурацию из .env
	cfg := config.Load()

	// Подключаемся к PostgreSQL
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database connected:", db != nil)

	// Создаём роутер
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	UserRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(UserRepository)
	AuthHandler := handler.NewAuthHandler(authService)
	// Auth routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/reg", AuthHandler.Register)
	})

	log.Println("server started on :8000")

	// Запускаем сервер
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}