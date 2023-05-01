package model

type OrderDetails struct {
	Id                 int64   `json:"id"`
	OrderId            int64   `json:"orderId"`
	ProductId          int64   `json:"productId"`
	ProductName        string  `json:"productName"`
	ProductDescription *string `json:"productDescription"`
	Price              float64 `json:"price"`
	Quantity           int64   `json:"quantity"`
}
