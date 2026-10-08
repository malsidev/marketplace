package routers

import (
	"net/http"
	"order/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	orderRepository := repository.NewOrderRepository(db)
	orderRouter := NewRouterOrder(orderRepository)

	r.Route("/", func(r chi.Router) {
		r.Mount("/order", orderRouter.OrderRouters())
	})

	return r
}
