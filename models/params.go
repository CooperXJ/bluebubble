package models

type SignUpParam struct {
	UserName   string `json:"username" binding:"required"`
	PassWord   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=PassWord"`
}

type LoginParam struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}
