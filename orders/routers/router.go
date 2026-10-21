package routers

import (
	"net/http"
	"order/repo"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	orderRepo := repo.NewOrderRepo(db)
	orderRouter := NewRouterOrder(orderRepo)

	r.Mount("/orders", orderRouter.orderRouters())

	return r
}
