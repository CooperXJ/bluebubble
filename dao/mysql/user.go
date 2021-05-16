package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"web_app/models"
)

const salt = "Cooper"

func InsertUser(user *models.User) (err error) {
	sqlStr:= "insert into user(user_id,username,password)values(?,?,?)"

	_, err = db.Exec(sqlStr, user.UserId, user.UserName, encodePassword(user.PassWord))
	if err != nil {
		return err
	}
	return nil
}

func QueryUserByName(username string) (flag bool ,err error) {
	sqlStr:="select count(*) from user where username=?"

	var count int

	err = db.Get(&count, sqlStr, username)

	if err != nil {
		return false,err
	}

	return count>0,nil
}

func encodePassword(password string) string {
	md := md5.New()
	md.Write([]byte(salt))
	return hex.EncodeToString(md.Sum([]byte(password)))
}