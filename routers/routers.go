package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller"
	"web_app/logger"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //如果是发布模式，设置为发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//r:=gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	r.POST("/signup", controller.SignUpHandler)

	r.POST("/login", controller.LoginHandler)
	return r
}
