package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

const salt = "Cooper"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

func InsertUser(user *models.User) (err error) {
	sqlStr := "insert into user(user_id,username,password)values(?,?,?)"

	_, err = db.Exec(sqlStr, user.UserId, user.UserName, encodePassword(user.PassWord))
	if err != nil {
		return err
	}
	return nil
}

func QueryUserByName(username string) (err error) {
	sqlStr := "select count(*) from user where username=?"

	var count int

	err = db.Get(&count, sqlStr, username)

	if err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}

	return nil
}

func FindUserByName(p *models.LoginParam) (err error) {
	var user = new(models.User)
	inputPassword := encodePassword(p.PassWord)
	sqlStr := "select user_id,username,password from user where username = ?"
	err = db.Get(user, sqlStr, inputPassword)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		//查询数据库失败
		return err
	}

	if inputPassword != user.PassWord {
		return ErrorInvalidPassword
	}
	return nil
}

func encodePassword(password string) string {
	md := md5.New()
	md.Write([]byte(salt))
	return hex.EncodeToString(md.Sum([]byte(password)))
}
