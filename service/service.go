package service

import (
	"github.com/fazarrahman/cognotiv/domain/order"
	"github.com/fazarrahman/cognotiv/domain/user"
	"github.com/fazarrahman/cognotiv/error"
	"github.com/fazarrahman/cognotiv/model"
	"github.com/gin-gonic/gin"
)

type Svc struct {
	OrderRepository order.OrderRepository
	UserRepository  user.UserRepository
}

// New ...
func New(_orderRepo order.OrderRepository, _userRepo user.UserRepository) *Svc {
	return &Svc{OrderRepository: _orderRepo, UserRepository: _userRepo}
}

// Service ...
type ServiceInterface interface {
	CreateOrder(ctx *gin.Context, r *model.Orders, username string) *error.Error
	GetOrderListByUserID(ctx *gin.Context, username string) ([]*model.Orders, *error.Error)
}
