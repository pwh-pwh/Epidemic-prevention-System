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
	r.GET("/userInfo", middlewares.JwtAuth(""), controller.UserInfo)
	//news 相关接口
	r.GET("/news", middlewares.JwtAuth(""), controller.News)
	r.GET("/chinaData", middlewares.JwtAuth(""), controller.ChinaData)
	r.GET("/riskarea", middlewares.JwtAuth(""), controller.GetRiskArea)
	r.GET("/history", middlewares.JwtAuth(""), controller.GetHistory)
	r.GET("/infiniteNews", middlewares.JwtAuth(""), controller.InfiniteNews)
	r.GET("/vaccineTopData", middlewares.JwtAuth(""), controller.VaccineTopData)
	r.GET("/chinaVaccineTrendData", middlewares.JwtAuth(""), controller.ChinaVaccineTrendData)
	r.GET("/rumor", middlewares.JwtAuth(""), controller.Rumor)

	r.POST("/upload", middlewares.JwtAuth(""), controller.Upload)
	//注册路由
	menuGroup := r.Group("/sys/menu")
	menuGroup.GET("/nav", middlewares.JwtAuth(""), controller.Nav)
	//公告路由
	noticeGroup := r.Group("/sys/notice")
	noticeGroup.GET("", middlewares.JwtAuth(""), controller.GetNotice)
	//未回归路由
	arGroup := r.Group("/access/return")
	arGroup.GET("/list", middlewares.JwtAuth("access:return:list"), controller.AccessReturnList)
	//出入登记路由
	aRgistGroup := r.Group("/access/register")
	aRgistGroup.GET("/list", middlewares.JwtAuth("access:register:list"), controller.GetAccessRegisterList)
	aRgistGroup.POST("", middlewares.JwtAuth("access:register:save"), controller.SaveAccessRegister)
	//goodsinfo 路由
	goodsInfoGroup := r.Group("/good/info")
	goodsInfoGroup.GET("/all", middlewares.JwtAuth("good:info:list"), controller.GetAllGoodsInfo)
	goodsInfoGroup.GET("/total", middlewares.JwtAuth("good:info:list"), controller.GetTotalGoodsInfo)
	goodsInfoGroup.GET("/list", middlewares.JwtAuth("good:info:list"), controller.GetListGoodsInfo)
	goodsInfoGroup.POST("", middlewares.JwtAuth("good:info:save"), controller.SaveGoodsInfo)
	goodsInfoGroup.PUT("", middlewares.JwtAuth("good:info:update"), controller.UpdateGoodsInfo)
	goodsInfoGroup.DELETE("", middlewares.JwtAuth("good:info:delete"), controller.DeleteGoodsInfo)
	//goodstype 路由
	goodsTypeGroup := r.Group("/good/type")
	goodsTypeGroup.GET("", middlewares.JwtAuth("good:type:list"), controller.GetSimpleListGoodsType)
	goodsTypeGroup.GET("/list", middlewares.JwtAuth("good:type:list"), controller.GetListGoodsType)
	goodsTypeGroup.POST("", middlewares.JwtAuth("good:type:save"), controller.SaveGoodsType)
	goodsTypeGroup.PUT("", middlewares.JwtAuth("good:type:update"), controller.UpdateGoodsType)
	goodsTypeGroup.DELETE("", middlewares.JwtAuth("good:type:delete"), controller.DeleteGoodsType)
	//物资出入库路由
	goodsStockGroup := r.Group("/good/stock")
	goodsStockGroup.GET("/list", middlewares.JwtAuth("good:stock:list"), controller.GetListGoodsStock)
	//good:stock:operate
	goodsStockGroup.POST("", middlewares.JwtAuth("good:stock:operate"), controller.SaveGoodsStock)
	///health/clock
	healthClockGroup := r.Group("/health/clock")
	healthClockGroup.GET("/list", middlewares.JwtAuth("health:clock:list"), controller.GetListHealthClock)
	healthClockGroup.GET("", middlewares.JwtAuth(""), controller.CheckHealthClock)
	healthClockGroup.POST("", middlewares.JwtAuth("health:clock:save"), controller.SaveHealthClock)
	// dept route
	deptGroup := r.Group("/sys/dept")
	deptGroup.GET("/list/:flag", controller.ListDept)
	// health/report route
	healthReportGroup := r.Group("/health/report")
	healthReportGroup.GET("", middlewares.JwtAuth(""), controller.CheckHealthReport)
	healthReportGroup.GET("/list", middlewares.JwtAuth("health:report:list"), controller.ListHealthReport)
	healthReportGroup.POST("", middlewares.JwtAuth("health:report:save"), controller.SaveHealthReport)
	// leave/apply route
	leaveApplyGroup := r.Group("/leave/apply")
	leaveApplyGroup.GET("/list", middlewares.JwtAuth("leave:apply:list,leave:record:list"), controller.GetListLeaveApply)
	leaveApplyGroup.POST("", middlewares.JwtAuth("leave:apply:save"), controller.SaveLeaveApply)
	leaveApplyGroup.PUT("", middlewares.JwtAuth("leave:apply:update,leave:record:examine"), controller.UpdateLeaveApply)
	// /monitor/redis
	monitorGroup := r.Group("/monitor")
	monitorGroup.GET("/redis", middlewares.JwtAuth("monitor:redis:list"), controller.GetRedisInfo)

	// /register
	registerGroup := r.Group("/register")
	registerGroup.POST("", middlewares.CaptchaMiddleware(), controller.Register)
	registerGroup.GET("/deptList", controller.DeptList)
	return r
}
