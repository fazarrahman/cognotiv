package service

import (
	"github.com/fazarrahman/cognotiv/domain/order"
	"github.com/fazarrahman/cognotiv/error"
	"github.com/fazarrahman/cognotiv/model"
	"github.com/gin-gonic/gin"
)

type Svc struct {
	OrderRepository order.OrderRepository
}

// New ...
func New(_orderRepo order.OrderRepository) *Svc {
	return &Svc{OrderRepository: _orderRepo}
}

// Service ...
type ServiceInterface interface {
	CreateOrder(ctx *gin.Context, r *model.Orders) *error.Error
}
