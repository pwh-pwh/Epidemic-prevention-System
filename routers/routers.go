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
	r.GET("userInfo", middlewares.JwtAuth(""), controller.UserInfo)
	r.GET("/news", middlewares.JwtAuth(""), controller.News)
	r.GET("/chinaData", middlewares.JwtAuth(""), controller.ChinaData)
	//注册路由
	menuGroup := r.Group("/sys/menu")
	menuGroup.GET("/nav", middlewares.JwtAuth(""), controller.Nav)
	apiGroup := r.Group("/api/v1")
	apiGroup.GET("/arlist", controller.GetAccessRegisterList)
	//公告路由
	noticeGroup := r.Group("/sys/notice")
	noticeGroup.GET("", middlewares.JwtAuth(""), controller.GetNotice)
	//未回归路由
	arGroup := r.Group("/access/return")
	arGroup.GET("/list", middlewares.JwtAuth("access:return:list"), controller.AccessReturnList)
	//出入登记路由
	aRgistGroup := r.Group("/access/register")
	aRgistGroup.GET("/list", controller.GetAccessRegisterList)
	return r
}
