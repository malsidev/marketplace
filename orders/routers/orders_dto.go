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
