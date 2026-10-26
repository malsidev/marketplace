package routers

type Order struct {
	UserID     int          `json:"UserID"`
	OrderItems []OrderItems `json:"OrderItems"`
}

type OrderItems struct {
	ProductID int `json:"ProductID"`
	Quantity  int `json:"Quantity"`
}

type ProductID struct {
	ProductID []int `json:"product_id"`
}

type RespProductData struct {
	ProductID int     `json:"ProductID"`
	Price     float64 `json:"Price"`
}

type OrderGet struct {
	UserID int `json:"UserID"`
}

type OrderResponse struct {
	ID         uint    `json:"id"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`
	ProductID  int     `json:"product_id"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}
