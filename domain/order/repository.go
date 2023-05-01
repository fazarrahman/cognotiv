package order

import (
	"context"

	entity "github.com/fazarrahman/cognotiv/domain/order/entity"
	"github.com/fazarrahman/cognotiv/error"
)

type OrderRepository interface {
	InsertOrder(ctx context.Context, order entity.Orders, details []entity.OrderDetails) *error.Error
	GetOrderList(ctx context.Context, userId *int64) ([]*entity.Orders, *error.Error)
	GetOrderDetailList(ctx context.Context, orderIds []int64) ([]*entity.OrderDetails, *error.Error)
	GetProductList(ctx context.Context, productIds []int64) ([]*entity.Products, *error.Error)
}
