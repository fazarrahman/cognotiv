package mysqldb

import (
	"context"
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
		`INSERT INTO orders (customer_id, order_date, status, created_at)
		VALUES(?, ?, ?, ?)`, order.CustomerId, now, 1, now)
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
