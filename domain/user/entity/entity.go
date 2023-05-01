package entity

import "time"

type Users struct {
	ID          int64     `db:"id"`
	Username    string    `db:"username"`
	Password    []byte    `db:"password"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	RoleId      int64     `db:"role_id"`
	Address     string    `db:"address"`
	Phone       string    `db:"phone"`
	Email       string    `db:"email"`
	CreatedDate time.Time `db:"created_at"`
}
