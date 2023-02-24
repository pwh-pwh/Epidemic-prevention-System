package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/user_service"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"log"
)

func LoginHander(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	//判断账号密码
	sysUser := user_service.GetByUserName(username)
	//登录成功
	if sysUser != nil && utils.Md5Crypt(password) == sysUser.Password {
		jwt, err := common.ReleaseToken(username)
		if err != nil {
			log.Printf("jwt gen error:%v", err)
		}
		context.Header(common.GetJwtHeader(), jwt)
		response.Success(context, "登录成功")
		return
		// TODO 插入sys_login_info 成功日志
	}
	//登录失败
	response.Fail(context, "用户名或密码错误")
	// TODO 插入sys_login_info
}
