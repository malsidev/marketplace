package routers

import (
	"encoding/json"
	"net/http"
	"order/repo"

	"github.com/go-chi/chi/v5"
)

type RouterOrders struct {
	orderRepo *repo.OrderRepo
}

func NewRouterOrder(orderRepo *repo.OrderRepo) *RouterOrders {
	return &RouterOrders{
		orderRepo: orderRepo,
	}
}

func (ro *RouterOrders) orderRouters() http.Handler {
	r := chi.NewRouter()

	r.Get("/", ro.orderGet)

	return r
}

func (ro *RouterOrders) orderGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := ro.orderRepo.GetOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := make([]OrderProductResponse, 0, len(items))

	for _, i := range items {
		resp = append(resp, OrderProductResponse{
			ID:        i.ID,
			OrderID:   i.OrderID,
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			Price:     i.Price,
		})
	}
	json.NewEncoder(w).Encode(resp)
}
