package entity

import "time"

type Orders struct {
	Id         int64     `db:"id"`
	CustomerId int64     `db:"customer_id"`
	OrderDate  time.Time `db:"order_date"`
	Status     int64     `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type OrderDetails struct {
	Id        int64     `db:"id"`
	OrderId   int64     `db:"order_id"`
	ProductId int64     `db:"product_id"`
	Quantity  int64     `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
