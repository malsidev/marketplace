package routers

import (
	"bytes"
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
	r.Get("/", ro.orderGet)

	return r
}

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
		bytes.NewBuffer(jsonData),
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
	fmt.Println(body)
	var price []RespProductData
	if err := json.Unmarshal(body, &price); err != nil {
		log.Println("Ошибка декодирования цены:", err)
		return
	}

	for i, item := range data.OrderItems {
		orderItems = append(orderItems, models.OrdersItems{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     float64(price[i].Price),
		})
	}

	ro.orderRepo.PostOrder(&ordres, orderItems)
	json.NewEncoder(w).Encode(data)
}

func (ro *RouterOrders) orderGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data OrderGet
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, `{"message":"invalid JSON"}`, http.StatusBadRequest)
		return
	}

	orders, err := ro.orderRepo.GetProduct(data.UserID)
	if err != nil {
		http.Error(w, `{"message":"invalid JSON"}`, http.StatusBadRequest)
		return
	}

	var response []OrderResponse

	for _, order := range orders {
		for _, item := range order.OrderItem {
			response = append(response, OrderResponse{
				ID:         uint(order.ID),
				Status:     order.Status,
				TotalPrice: order.TotalPrice,
				ProductID:  item.ProductID,
				Price:      item.Price,
				Quantity:   item.Quantity,
			})
		}
	}

	var list_productID []int

	for _, i := range response {
		list_productID = append(list_productID, i.ProductID)
	}

	fmt.Println(list_productID)
	json.NewEncoder(w).Encode(response)
}
