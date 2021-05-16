package services

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.SignUpParam) (err error) {
	//查询用户存在与否
	err = mysql.QueryUserByName(p.UserName)
	if err != nil {
		return err
	}

	//生成唯一Id
	userId := snowflake.GenID()

	user := &models.User{
		UserName: p.UserName,
		PassWord: p.PassWord,
		UserId:   userId,
	}
	return mysql.InsertUser(user)
}

func Login(p *models.LoginParam) (err error) {
	return mysql.FindUserByName(p)
}
