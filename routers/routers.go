package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/controller"
	_ "github.com/pwh-pwh/Epidemic-prevention-System/docs"
	"github.com/pwh-pwh/Epidemic-prevention-System/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}
	r := gin.New()
	r.Use(middlewares.Recover)
	r.Use(middlewares.CommonLogInterceptor())
	r.Use(middlewares.CorsMiddleware())
	r.GET("/captcha", controller.GetCaptcha)
	r.POST("/login", middlewares.CaptchaMiddleware(), controller.LoginHander)
	r.GET("/userInfo", middlewares.JwtAuth(""), controller.UserInfo)
	r.POST("/upload", middlewares.JwtAuth(""), controller.Upload)
	r.POST("/logout", controller.Logout)
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//news 相关接口
	setUpNews(r)
	//注册路由
	setUpMenu(r)
	//公告路由
	setUpNotice(r)
	//未回归路由
	setUpAReturn(r)
	//出入登记路由
	setUpAr(r)
	//goodsinfo 路由
	setUpGoodsInfo(r)
	//goodstype 路由
	setUpGoodsType(r)
	//物资出入库路由
	setUpStock(r)
	///health/clock
	setUpHealthClock(r)
	// dept route
	setUpDept(r)
	// health/report route
	setUpHealthReport(r)
	// leave/apply route
	setUpLeaveApply(r)
	// /monitor/serve
	setUpServe(r)
	// /register
	setUpRegister(r)
	// /sys/loginInfo
	setUploginInfo(r)
	// /sys/operateLog
	setUpOpLog(r)
	// /sys/role
	setUpRole(r)
	// /sys/user
	setUpUser(r)
	return r
}
func setUpAReturn(r *gin.Engine) {
	arGroup := r.Group("/access/return")
	arGroup.GET("/list", middlewares.JwtAuth("access:return:list"), controller.AccessReturnList)
}
func setUpAr(r *gin.Engine) {
	aRgistGroup := r.Group("/access/register")
	aRgistGroup.GET("/list", middlewares.JwtAuth("access:register:list"), controller.GetAccessRegisterList)
	aRgistGroup.POST("", middlewares.JwtAuth("access:register:save"),
		middlewares.NewMetaHandler().SetTitle("出入登记").SetType("添加记录").SetMethod("SaveAccessRegister").ToHFunc(), controller.SaveAccessRegister)
}
func setUpGoodsInfo(r *gin.Engine) {
	goodsInfoGroup := r.Group("/good/info", middlewares.NewMetaHandler().SetTitle("物资资料").ToHFunc())
	goodsInfoGroup.GET("/all", middlewares.JwtAuth("good:info:list"), controller.GetAllGoodsInfo)
	goodsInfoGroup.GET("/total", middlewares.JwtAuth("good:info:list"), controller.GetTotalGoodsInfo)
	goodsInfoGroup.GET("/list", middlewares.JwtAuth("good:info:list"), controller.GetListGoodsInfo)
	goodsInfoGroup.POST("", middlewares.JwtAuth("good:info:save"), middlewares.NewMetaHandler().SetType("添加物资").SetMethod("SaveGoodsInfo").ToHFunc(), controller.SaveGoodsInfo)
	goodsInfoGroup.PUT("", middlewares.JwtAuth("good:info:update"), middlewares.NewMetaHandler().SetType("修改物资").SetMethod("UpdateGoodsInfo").ToHFunc(), controller.UpdateGoodsInfo)
	goodsInfoGroup.DELETE("", middlewares.JwtAuth("good:info:delete"), middlewares.NewMetaHandler().SetType("删除物资").SetMethod("DeleteGoodsInfo").ToHFunc(), controller.DeleteGoodsInfo)
}
func setUpGoodsType(r *gin.Engine) {
	goodsTypeGroup := r.Group("/good/type", middlewares.NewMetaHandler().SetTitle("物资类型").ToHFunc())
	goodsTypeGroup.GET("", middlewares.JwtAuth("good:type:list"), controller.GetSimpleListGoodsType)
	goodsTypeGroup.GET("/list", middlewares.JwtAuth("good:type:list"), controller.GetListGoodsType)
	goodsTypeGroup.POST("", middlewares.JwtAuth("good:type:save"), middlewares.NewMetaHandler().SetType("添加物资类型").SetMethod("SaveGoodsType").ToHFunc(), controller.SaveGoodsType)
	goodsTypeGroup.PUT("", middlewares.JwtAuth("good:type:update"), middlewares.NewMetaHandler().SetType("更新物资类型").SetMethod("UpdateGoodsType").ToHFunc(), controller.UpdateGoodsType)
	goodsTypeGroup.DELETE("", middlewares.JwtAuth("good:type:delete"), middlewares.NewMetaHandler().SetType("删除物资类型").SetMethod("DeleteGoodsType").ToHFunc(), controller.DeleteGoodsType)
}
func setUpStock(r *gin.Engine) {
	goodsStockGroup := r.Group("/good/stock")
	goodsStockGroup.GET("/list", middlewares.JwtAuth("good:stock:list"), controller.GetListGoodsStock)
	goodsStockGroup.POST("", middlewares.JwtAuth("good:stock:operate"),
		middlewares.NewMetaHandler().SetTitle("物资库存").SetType("物资出入库").SetMethod("SaveGoodsStock").ToHFunc(), controller.SaveGoodsStock)
}
func setUpHealthClock(r *gin.Engine) {
	healthClockGroup := r.Group("/health/clock")
	healthClockGroup.GET("/list", middlewares.JwtAuth("health:clock:list"), controller.GetListHealthClock)
	healthClockGroup.GET("", middlewares.JwtAuth(""), controller.CheckHealthClock)
	healthClockGroup.POST("", middlewares.JwtAuth("health:clock:save"),
		middlewares.NewMetaHandler().SetTitle("健康打卡").SetType("添加打卡").SetMethod("SaveHealthClock").ToHFunc(), controller.SaveHealthClock)
}

func setUpDept(r *gin.Engine) {
	deptGroup := r.Group("/sys/dept", middlewares.NewMetaHandler().SetTitle("部门管理").ToHFunc())
	deptGroup.GET("/list/:flag", controller.ListDept)
	deptGroup.DELETE("/:id", middlewares.JwtAuth("sys:dept:delete"),
		middlewares.NewMetaHandler().SetType("删除部门").SetMethod("DeleteDeptById").ToHFunc(), controller.DeleteDeptById)
	deptGroup.PUT("", middlewares.JwtAuth("sys:dept:update"),
		middlewares.NewMetaHandler().SetType("修改部门").SetMethod("UpdateDept").ToHFunc(), controller.UpdateDept)
	deptGroup.POST("", middlewares.JwtAuth("sys:dept:save"),
		middlewares.NewMetaHandler().SetType("新建部门").SetMethod("SaveDept").ToHFunc(), controller.SaveDept)
}

func setUpHealthReport(r *gin.Engine) {
	healthReportGroup := r.Group("/health/report")
	healthReportGroup.GET("", middlewares.JwtAuth(""), controller.CheckHealthReport)
	healthReportGroup.GET("/list", middlewares.JwtAuth("health:report:list"), controller.ListHealthReport)
	healthReportGroup.POST("", middlewares.JwtAuth("health:report:save"),
		middlewares.NewMetaHandler().SetTitle("健康报告").SetType("添加报告").SetMethod("SaveHealthReport").ToHFunc(), controller.SaveHealthReport)
}

func setUpLeaveApply(r *gin.Engine) {
	leaveApplyGroup := r.Group("/leave/apply", middlewares.NewMetaHandler().SetTitle("请假管理").ToHFunc())
	leaveApplyGroup.GET("/list", middlewares.JwtAuth("leave:apply:list,leave:record:list"), controller.GetListLeaveApply)
	leaveApplyGroup.POST("", middlewares.JwtAuth("leave:apply:save"),
		middlewares.NewMetaHandler().SetType("新增请假").SetMethod("SaveLeaveApply").ToHFunc(), controller.SaveLeaveApply)
	leaveApplyGroup.PUT("", middlewares.JwtAuth("leave:apply:update,leave:record:examine"),
		middlewares.NewMetaHandler().SetType("修改请假").SetMethod("UpdateLeaveApply").ToHFunc(), controller.UpdateLeaveApply)
}

func setUpServe(r *gin.Engine) {
	monitorGroup := r.Group("/monitor")
	monitorGroup.GET("/server", middlewares.JwtAuth("monitor:server:list"), controller.ServerInfo)
	monitorGroup.GET("/redis", middlewares.JwtAuth("monitor:redis:list"), controller.GetRedisInfo)
}

func setUpRegister(r *gin.Engine) {
	registerGroup := r.Group("/register")
	registerGroup.POST("", middlewares.CaptchaMiddleware(), controller.Register)
	registerGroup.GET("/deptList", controller.DeptList)
}

func setUploginInfo(r *gin.Engine) {
	loginInfoGroup := r.Group("/sys/loginInfo", middlewares.NewMetaHandler().SetTitle("登录日志管理").ToHFunc())
	loginInfoGroup.POST("", middlewares.JwtAuth("sys:login:clear"),
		middlewares.NewMetaHandler().SetType("清空登录日志").SetMethod("ClearLoginInfo").ToHFunc(), controller.ClearLoginInfo)
	loginInfoGroup.DELETE("", middlewares.JwtAuth("sys:login:delete"),
		middlewares.NewMetaHandler().SetType("删除登录日志").SetMethod("DeleteLoginInfo").ToHFunc(), controller.DeleteLoginInfo)
	loginInfoGroup.GET("/list", middlewares.JwtAuth("sys:login:list"), controller.ListLoginInfo)
}

func setUpOpLog(r *gin.Engine) {
	opLogGroup := r.Group("/sys/operateLog", middlewares.NewMetaHandler().SetTitle("操作日志管理").ToHFunc())
	opLogGroup.POST("", middlewares.JwtAuth("monitor:operate:clear"),
		middlewares.NewMetaHandler().SetType("清除日志").SetMethod("ClearOpLog").ToHFunc(), controller.ClearOpLog)
	opLogGroup.GET("/list", middlewares.JwtAuth("monitor:operate:list"), controller.ListOpLog)
	opLogGroup.DELETE("", middlewares.JwtAuth("monitor:operate:delete"),
		middlewares.NewMetaHandler().SetType("删除日志").SetMethod("DeleteOpLog").ToHFunc(), controller.DeleteOpLog)
}

func setUpRole(r *gin.Engine) {
	roleGroup := r.Group("/sys/role", middlewares.NewMetaHandler().SetTitle("角色管理").ToHFunc())
	roleGroup.DELETE("", middlewares.JwtAuth("sys:role:delete"),
		middlewares.NewMetaHandler().SetType("删除角色").SetMethod("DeleteRole").ToHFunc(), controller.DeleteRole)
	roleGroup.GET("/info/:id", middlewares.JwtAuth(""), controller.InfoRole)
	roleGroup.GET("/list", middlewares.JwtAuth("sys:role:list"), controller.ListRole)
	roleGroup.POST("", middlewares.JwtAuth("sys:role:save"),
		middlewares.NewMetaHandler().SetType("添加角色").SetMethod("AddRole").ToHFunc(), controller.AddRole)
	roleGroup.PUT("", middlewares.JwtAuth("sys:role:update"),
		middlewares.NewMetaHandler().SetType("修改角色").SetMethod("EditRole").ToHFunc(), controller.EditRole)
}

func setUpUser(r *gin.Engine) {
	userGroup := r.Group("/sys/user", middlewares.NewMetaHandler().SetTitle("用户管理").ToHFunc())
	userGroup.POST("/avatar", middlewares.JwtAuth(""), controller.Avatar)
	userGroup.GET("/updatePassword", middlewares.JwtAuth("sys:user:update"),
		middlewares.NewMetaHandler().SetType("更新密码").SetMethod("UpdatePassword").ToHFunc(), controller.UpdatePassword)
	userGroup.POST("/updateInfo", middlewares.JwtAuth("sys:user:update"),
		middlewares.NewMetaHandler().SetType("更新用户信息").SetMethod("UpdateInfo").ToHFunc(), controller.UpdateInfo)
	userGroup.POST("/reset", middlewares.JwtAuth("sys:user:repass"),
		middlewares.NewMetaHandler().SetType("重置密码").SetMethod("ResetPwd").ToHFunc(), controller.ResetPwd)
	userGroup.POST("/userRole/:id", middlewares.JwtAuth("sys:user:role"),
		middlewares.NewMetaHandler().SetType("分配用户角色").SetMethod("ApplyUserRole").ToHFunc(), controller.ApplyUserRole)
	userGroup.DELETE("", middlewares.JwtAuth("sys:user:delete"),
		middlewares.NewMetaHandler().SetType("删除用户").SetMethod("DeleteUesr").ToHFunc(), controller.DeleteUesr)
	userGroup.PUT("", middlewares.JwtAuth("sys:user:update"),
		middlewares.NewMetaHandler().SetType("更新用户").SetMethod("UpdateUser").ToHFunc(), controller.UpdateUser)
	userGroup.POST("", middlewares.JwtAuth("sys:user:save"),
		middlewares.NewMetaHandler().SetType("添加用户").SetMethod("AddUser").ToHFunc(), controller.AddUser)
	userGroup.GET("/list", middlewares.JwtAuth("sys:user:list"), controller.ListUser)
	userGroup.GET("/info/:id", controller.Info)
}

func setUpNotice(r *gin.Engine) {
	noticeGroup := r.Group("/sys/notice", middlewares.NewMetaHandler().SetTitle("公告管理").ToHFunc())
	noticeGroup.GET("", middlewares.JwtAuth(""), controller.GetNotice)
	noticeGroup.GET("/list", middlewares.JwtAuth("monitor:notice:list"), controller.ListNotice)
	noticeGroup.POST("", middlewares.JwtAuth("monitor:notice:save"),
		middlewares.NewMetaHandler().SetType("新建公告").SetMethod("SaveNotice").ToHFunc(), controller.SaveNotice)
	noticeGroup.PUT("", middlewares.JwtAuth("monitor:notice:update"),
		middlewares.NewMetaHandler().SetType("更新公告").SetMethod("UpdateNotice").ToHFunc(), controller.UpdateNotice)
	noticeGroup.DELETE("", middlewares.JwtAuth("monitor:notice:delete"),
		middlewares.NewMetaHandler().SetType("删除公告").SetMethod("DeleteNotice").ToHFunc(), controller.DeleteNotice)
	noticeGroup.GET("/:id", middlewares.JwtAuth("monitor:notice:set"), controller.SetNotice)
}

func setUpNews(r *gin.Engine) {
	r.GET("/news", middlewares.JwtAuth(""), controller.News)
	r.GET("/chinaData", middlewares.JwtAuth(""), controller.ChinaData)
	r.GET("/riskarea", middlewares.JwtAuth(""), controller.GetRiskArea)
	r.GET("/history", middlewares.JwtAuth(""), controller.GetHistory)
	r.GET("/infiniteNews", middlewares.JwtAuth(""), controller.InfiniteNews)
	r.GET("/vaccineTopData", middlewares.JwtAuth(""), controller.VaccineTopData)
	r.GET("/chinaVaccineTrendData", middlewares.JwtAuth(""), controller.ChinaVaccineTrendData)
	r.GET("/rumor", middlewares.JwtAuth(""), controller.Rumor)
}

func setUpMenu(r *gin.Engine) {
	menuGroup := r.Group("/sys/menu", middlewares.NewMetaHandler().SetTitle("菜单管理").ToHFunc())
	menuGroup.GET("/nav", middlewares.JwtAuth(""), controller.Nav)
	menuGroup.GET("/list", middlewares.JwtAuth("sys:menu:list"), controller.ListMenu)
	menuGroup.DELETE("/:id", middlewares.JwtAuth("sys:menu:delete"),
		middlewares.NewMetaHandler().SetType("删除菜单").SetMethod("DeleteMenu").ToHFunc(), controller.DeleteMenu)
	menuGroup.POST("", middlewares.JwtAuth("sys:menu:save"),
		middlewares.NewMetaHandler().SetType("新建菜单").SetMethod("SaveMenu").ToHFunc(), controller.SaveMenu)
	menuGroup.PUT("", middlewares.JwtAuth("sys:menu:update"),
		middlewares.NewMetaHandler().SetType("修改菜单").SetMethod("EditMenu").ToHFunc(), controller.EditMenu)
}
