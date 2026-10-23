package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"order/models"
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

	// r.Get("/", ro.orderGet)
	r.Post("/", ro.orderPost)

	return r
}

// func (ro *RouterOrders) orderGet(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	items, err := ro.orderRepo.GetOrder()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	resp := make([]OrderProductResponse, 0, len(items))

// 	for _, i := range items {
// 		resp = append(resp, Orderfunc (r *OrderRepo) PostOrder(order *models.Orders, items *models.OrdersItems) error {
// 			ID:        i.ID,
// 			OrderID:   i.OrderID,
// 			ProductID: i.ProductID,
// 			Quantity:  i.Quantity,
// 			Price:     i.Price,
// 		})
// 	}
// 	json.NewEncoder(w).Encode(resp)
// }

func (ro *RouterOrders) orderPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data Order
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, `{"message":"invalid JSON"}`, http.StatusBadRequest)
		return
	}

	id := 1
	curl := fmt.Sprintf("http://127.0.0.1:8001/products/%d", id)
	resp, err := http.Get(curl)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения ответа:", err)
		return
	}

	var price float64
	if err := json.Unmarshal(body, &price); err != nil {
		log.Println("Ошибка декодирования цены:", err)
		return
	}

	ordres := models.Orders{
		UserID:     data.UserID,
		TotalPrice: int(price),
	}

	orderItems := models.OrdersItems{
		ProductID: data.OrderItems[0].ProductID,
		Quantity:  data.OrderItems[0].Quantity,
	}

	ro.orderRepo.PostOrder(&ordres, &orderItems)

	log.Println("Статус:", resp.Status)
	log.Println("Ответ:", string(body)) /* Ответ: 65000.0 */
	json.NewEncoder(w).Encode(data)
}
