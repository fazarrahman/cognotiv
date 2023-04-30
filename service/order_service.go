package service

import (
	ent "github.com/fazarrahman/cognotiv/domain/order/entity"
	"github.com/fazarrahman/cognotiv/error"
	"github.com/fazarrahman/cognotiv/model"
	"github.com/gin-gonic/gin"
)

func (s *Svc) CreateOrder(ctx *gin.Context, r *model.Orders) *error.Error {
	if r.CustomerId == 0 {
		return error.BadRequest("Customer id is required")
	} else if len(r.Details) == 0 {
		return error.BadRequest("Order detail is required")
	}

	var order = ent.Orders{
		CustomerId: r.CustomerId,
	}

	var orderDetails []ent.OrderDetails
	for _, d := range r.Details {
		if d.ProductId == 0 {
			return error.BadRequest("Product id is required")
		} else if d.Quantity < 1 {
			return error.BadRequest("Quantity is required")
		}
		orderDetails = append(orderDetails, ent.OrderDetails{
			ProductId: d.ProductId,
			Quantity:  d.Quantity,
		})
	}
	err := s.OrderRepository.InsertOrder(ctx, order, orderDetails)
	if err != nil {
		return err
	}
	return nil
}
