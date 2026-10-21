package routers

type Order struct {
	OrderID int `json:"OrderID"`
}

type OrderProductResponse struct {
	ID        int `json:"OrderItemsID"`
	OrderID   int `json:"OrderID"`
	ProductID int `json:"ProductID"`
	Quantity  int `json:"Quantity"`
	Price     int `json:"Price"`
}
