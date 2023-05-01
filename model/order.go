package model

import "time"

type Orders struct {
	Id        int64          `json:"id"`
	UserId    int64          `json:"userId" validate:"required"`
	OrderDate *time.Time     `json:"orderDate"`
	Status    *int64         `json:"status"`
	Details   []OrderDetails `json:"details"`
}
