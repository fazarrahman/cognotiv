package entity

import "time"

type Orders struct {
	Id        int64      `db:"id"`
	UserId    int64      `db:"user_id"`
	OrderDate time.Time  `db:"order_date"`
	Status    int64      `db:"status"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type OrderDetails struct {
	Id        int64      `db:"id"`
	OrderId   int64      `db:"order_id"`
	ProductId int64      `db:"product_id"`
	Quantity  int64      `db:"quantity"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type Products struct {
	Id          int64      `db:"id"`
	Name        string     `db:"name"`
	Price       float64    `db:"price"`
	Description *string    `db:"description"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
