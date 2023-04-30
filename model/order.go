package model

import "time"

type Orders struct {
	Id         int64          `json:"id"`
	CustomerId int64          `json:"customerId" validate:"required"`
	OrderDate  *time.Time     `json:"orderDate"`
	Status     *int64         `json:"status"`
	Details    []OrderDetails `json:"details"`
}
