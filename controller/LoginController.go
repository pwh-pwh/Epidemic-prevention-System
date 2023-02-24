package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/user_service"
	"github.com/pwh-pwh/Epidemic-prevention-System/task"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"log"
	"time"
)

func LoginHander(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	//判断账号密码
	sysUser := user_service.GetByUserName(username)
	userAgent := user_agent.New(context.GetHeader("User-Agent"))
	browser, bV := userAgent.Browser()
	browser = browser + " " + bV
	loginInfo := models.SysLoginInfo{
		IP:        context.RemoteIP(),
		Username:  username,
		Browser:   browser,
		Os:        userAgent.OS(),
		Status:    1,
		Msg:       "登录成功",
		LoginTime: models.LocalTime(time.Now()),
	}
	//登录成功
	if sysUser != nil && utils.Md5Crypt(password) == sysUser.Password {
		jwt, err := common.ReleaseToken(username)
		if err != nil {
			log.Printf("jwt gen error:%v", err)
		}
		context.Header(common.GetJwtHeader(), jwt)
		response.Success(context, "登录成功")
		task.AddTask(func() {
			addr := utils.GetLocation(loginInfo.IP)
			loginInfo.LoginLocation = addr
			service.SaveLoginInfo(&loginInfo)
		})
		return
	}
	//登录失败
	response.Fail(context, "用户名或密码错误")
	// 插入sys_login_info
	loginInfo.Msg = "用户名或密码错误"
	loginInfo.Status = 0
	task.AddTask(func() {
		addr := utils.GetLocation(loginInfo.IP)
		loginInfo.LoginLocation = addr
		service.SaveLoginInfo(&loginInfo)
	})
}
