package mysqldb

import (
	"context"
	"time"

	ue "github.com/fazarrahman/cognotiv/domain/user/entity"
	"github.com/fazarrahman/cognotiv/error"
	"github.com/jmoiron/sqlx"
)

type Mysqldb struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Mysqldb {
	return &Mysqldb{db: db}
}

func (m *Mysqldb) GetUserByUsername(ctx context.Context, username string) (*ue.Users, *error.Error) {
	var user ue.Users
	err := m.db.GetContext(ctx, &user, `SELECT id, username, role_id, first_name, last_name, address, phone, email 
	FROM users u WHERE username = ?`, username)
	if err != nil {
		return nil, error.InternalServerError(err.Error())
	}
	return &user, nil
}

func (m *Mysqldb) GetUserByEmail(ctx context.Context, email string) (*ue.Users, *error.Error) {
	var user ue.Users
	err := m.db.GetContext(ctx, &user, `SELECT id, username, password, role_id, first_name, last_name, address, phone, email 
	FROM users u WHERE email = ?`, email)
	if err != nil {
		return nil, error.InternalServerError(err.Error())
	}
	return &user, nil
}

func (m *Mysqldb) GetRoleCodeByUsername(ctx context.Context, username string) (*string, *error.Error) {
	var roleCode string
	err := m.db.GetContext(ctx, &roleCode, `select r.code from roles r 
	inner join users u on r.id = u.role_id where username = ?`, username)
	if err != nil {
		return nil, error.InternalServerError(err.Error())
	}
	return &roleCode, nil
}

func (m *Mysqldb) InsertUser(ctx context.Context, user *ue.Users) *error.Error {
	user.CreatedDate = time.Now()
	_, err := m.db.NamedExecContext(ctx, `INSERT INTO users
	(username, password, role_id, first_name, last_name, address, phone, email, created_at)
	VALUES(:username, :password, :role_id, :first_name, :last_name, :address, :phone, :email, :created_at)`, user)

	if err != nil {
		return error.InternalServerError(err.Error())
	}

	return nil
}
