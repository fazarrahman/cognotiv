package model

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName"`
	RoleId    int64  `json:"roleId" validate:"required"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `jsondb:"email" validate:"required,email"`
}
