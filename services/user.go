package services

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.SignUpParam) (err error) {
	//查询用户存在与否
	exits, err := mysql.QueryUserByName(p.UserName)
	if err != nil {
		//查询数据库出错
		return err
	}

	if exits{
		return errors.New("该用户已存在")
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
