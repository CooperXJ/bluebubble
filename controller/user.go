package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/services"
)

func SignUpHandler(c *gin.Context) {

	p := new(models.SignUpParam)
	err := c.ShouldBind(p)
	if err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		//c.JSON(http.StatusOK,gin.H{
		//	"msg": err.Error(),
		//})
		return
	}

	err = services.SignUp(p)
	//if err != nil {
	//	c.JSON(http.StatusOK,gin.H{
	//		"msg": err.Error(),
	//	})
	//}
	if err != nil {
		zap.L().Error("sign failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	p := new(models.LoginParam)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}

	//登录
	err := services.Login(p)
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	ResponseSuccess(c, nil)
}
