package models

type User struct {
	UserName string `db:"username"`
	PassWord string	`db:"password"`
	UserId int64 `db:"user_id"`
}
