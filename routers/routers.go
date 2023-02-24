package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/controller"
	"github.com/pwh-pwh/Epidemic-prevention-System/middlewares"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}
	r := gin.New()
	r.Use(middlewares.Recover)
	r.Use(middlewares.CorsMiddleware())
	r.GET("/captcha", controller.GetCaptcha)
	r.POST("/login", middlewares.CaptchaMiddleware(), controller.LoginHander)
	//注册路由
	apiGroup := r.Group("/api/v1")
	apiGroup.GET("/arlist", controller.GetAccessRegisterList)
	return r
}
