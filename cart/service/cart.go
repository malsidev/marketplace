package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CartItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ProductsRequest struct {
	Items []CartItemRequest `json:"items"`
}

func GetProduct(items []CartItemRequest) ([]byte, error) {
	data := ProductsRequest{
		Items: items,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		"http://127.0.0.1:8001/cart",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	fmt.Println(string(body))

	return io.ReadAll(resp.Body)
}
