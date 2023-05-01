package service

import (
	ent "github.com/fazarrahman/cognotiv/domain/order/entity"
	"github.com/fazarrahman/cognotiv/error"
	"github.com/fazarrahman/cognotiv/model"
	"github.com/gin-gonic/gin"
)

const (
	ROLE_CODE_ADMIN    string = "ADMIN"
	ROLE_CODE_CUSTOMER string = "CUSTOMER"
)

func (s *Svc) CreateOrder(ctx *gin.Context, r *model.Orders, username string) *error.Error {
	if r.UserId == 0 {
		return error.BadRequest("User id is required")
	} else if len(r.Details) == 0 {
		return error.BadRequest("Order detail is required")
	}

	user, err := s.UserRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}
	var order = ent.Orders{
		UserId: user.ID,
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
	err = s.OrderRepository.InsertOrder(ctx, order, orderDetails)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) GetOrderListByUserID(ctx *gin.Context, username string) ([]*model.Orders, *error.Error) {
	if username == "" {
		return nil, error.Unauthorized("Invalid authorization token")
	}
	roleCode, err := s.UserRepository.GetRoleCodeByUsername(ctx, username)
	if roleCode == nil {
		return nil, error.BadRequest("Role code not found")
	}
	if err != nil {
		return nil, err
	}
	var orders []*ent.Orders
	if *roleCode == ROLE_CODE_ADMIN {
		orders, err = s.OrderRepository.GetOrderList(ctx, nil)
		if err != nil {
			return nil, err
		}
	} else {
		user, err := s.UserRepository.GetUserByUsername(ctx, username)
		if err != nil {
			return nil, err
		}
		orders, err = s.OrderRepository.GetOrderList(ctx, &user.ID)
		if err != nil {
			return nil, err
		}
	}

	if orders == nil {
		return nil, nil
	}

	var orderIds []int64
	for _, o := range orders {
		orderIds = append(orderIds, o.Id)
	}

	details, err := s.OrderRepository.GetOrderDetailList(ctx, orderIds)
	if err != nil {
		return nil, err
	}

	var detailMap = make(map[int64][]ent.OrderDetails)
	for _, d := range details {
		detailMap[d.OrderId] = append(detailMap[d.OrderId], *d)
	}

	var productIds []int64
	var prodIdMap = make(map[int64]bool)
	for _, d := range details {
		if !prodIdMap[d.ProductId] {
			productIds = append(productIds, d.ProductId)
			prodIdMap[d.ProductId] = true
		}
	}
	products, err := s.OrderRepository.GetProductList(ctx, productIds)
	if err != nil {
		return nil, err
	}

	var prodMap = make(map[int64]ent.Products)
	for _, p := range products {
		prodMap[p.Id] = *p
	}

	var orderModels []*model.Orders
	for _, o := range orders {
		orderModels = append(orderModels, &model.Orders{
			Id:        o.Id,
			UserId:    o.UserId,
			OrderDate: &o.OrderDate,
			Status:    &o.Status,
		})
	}

	for _, o := range orderModels {
		for _, d := range detailMap[o.Id] {
			o.Details = append(o.Details, model.OrderDetails{
				Id:                 d.Id,
				OrderId:            d.OrderId,
				ProductId:          d.ProductId,
				Quantity:           d.Quantity,
				ProductName:        prodMap[d.ProductId].Name,
				Price:              prodMap[d.ProductId].Price,
				ProductDescription: prodMap[d.ProductId].Description,
			})
		}
	}

	return orderModels, nil
}
