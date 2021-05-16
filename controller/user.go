package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"web_app/models"
	"web_app/services"
)

func SignUpHandler(c *gin.Context)  {

	p := new(models.SignUpParam)
	err := c.ShouldBind(p)
	if err != nil {
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		c.JSON(http.StatusOK,gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = services.SignUp(p)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"msg": err.Error(),
		})
	}

	//c.JSON(http.StatusOK,"ok")
}
