package model

type OrderDetails struct {
	Id        int64 `json:"id"`
	OrderId   int64 `json:"orderId"`
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}
