package routers

import (
	"bytes"
	"encoding/json"
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

	var orderItems []models.OrdersItems

	ordres := models.Orders{
		UserID: data.UserID,
	}

	var list_productID []int

	for _, item := range data.OrderItems {
		list_productID = append(list_productID, item.ProductID)
	}

	listID := ProductID{ProductID: list_productID}
	jsonData, err := json.Marshal(listID)
	if err != nil {
		log.Fatalln(err)
	}

	curl := "http://127.0.0.1:8001/products/id"
	resp, err := http.Post(
		curl,
		"application/json",
		bytes.Buffer(jsonData),
	)

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

	orderItems = append(orderItems, models.OrdersItems{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     int(price),
	})

	ro.orderRepo.PostOrder(&ordres, orderItems)
	json.NewEncoder(w).Encode(data)
}
