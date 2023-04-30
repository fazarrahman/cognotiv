package order

import (
	"context"

	entity "github.com/fazarrahman/cognotiv/domain/order/entity"
	"github.com/fazarrahman/cognotiv/error"
)

type OrderRepository interface {
	InsertOrder(ctx context.Context, order entity.Orders, details []entity.OrderDetails) *error.Error
}
