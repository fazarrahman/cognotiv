package mysqldb

import (
	"context"
	"log"
	"time"

	entity "github.com/fazarrahman/cognotiv/domain/order/entity"
	"github.com/fazarrahman/cognotiv/error"
	"github.com/jmoiron/sqlx"
)

type Mysqldb struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Mysqldb {
	return &Mysqldb{db: db}
}

func (m *Mysqldb) InsertOrder(ctx context.Context, order entity.Orders, details []entity.OrderDetails) *error.Error {
	now := time.Now()
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return error.InternalServerError(err.Error())
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(ctx,
		`INSERT INTO orders (user_id, order_date, status, created_at)
		VALUES(?, ?, ?, ?)`, order.UserId, now, 1, now)
	if err != nil {
		return error.InternalServerError(err.Error())
	}
	orderId, err := res.LastInsertId()
	if err != nil {
		return error.InternalServerError(err.Error())
	}

	for _, d := range details {
		_, err := tx.ExecContext(ctx, `INSERT INTO order_details
		(order_id, product_id, quantity, created_at) 
		VALUES(?, ?, ?, ?)`, orderId, d.ProductId, d.Quantity, now)
		if err != nil {
			return error.InternalServerError(err.Error())
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return error.InternalServerError(err.Error())
	}

	return nil
}

func (m *Mysqldb) GetOrderList(ctx context.Context, userId *int64) ([]*entity.Orders, *error.Error) {
	var orders []*entity.Orders
	if userId == nil {
		err := m.db.SelectContext(ctx, &orders, `SELECT id, user_id, order_date, status, created_at, updated_at
		FROM orders where status = 1`)
		if err != nil {
			return nil, error.InternalServerError(err.Error())
		}
	} else {
		err := m.db.SelectContext(ctx, &orders, `SELECT id, user_id, order_date, status, created_at, updated_at
		FROM orders where user_id = ? AND status = 1`, &userId)
		if err != nil {
			return nil, error.InternalServerError(err.Error())
		}
	}

	return orders, nil
}

func (m *Mysqldb) GetOrderDetailList(ctx context.Context, orderIds []int64) ([]*entity.OrderDetails, *error.Error) {
	query, args, err := sqlx.In("SELECT id, order_id, product_id, quantity from order_details od where order_id IN (?)", orderIds)
	if err != nil {
		log.Fatal(err)
	}

	var orderDetails []*entity.OrderDetails
	query = m.db.Rebind(query)
	err = m.db.SelectContext(ctx, &orderDetails, query, args...)
	if err != nil {
		return nil, error.InternalServerError(err.Error())
	}
	return orderDetails, nil
}

func (m *Mysqldb) GetProductList(ctx context.Context, productIds []int64) ([]*entity.Products, *error.Error) {
	query, args, err := sqlx.In(`SELECT id, name, price, description FROM products where id IN (?)`, productIds)
	if err != nil {
		log.Fatal(err)
	}

	var products []*entity.Products
	query = m.db.Rebind(query)
	err = m.db.SelectContext(ctx, &products, query, args...)
	if err != nil {
		return nil, error.InternalServerError(err.Error())
	}
	return products, nil
}
