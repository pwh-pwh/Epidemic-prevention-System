package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/user_service"
	"strconv"
)

func Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	roleTypeS := ctx.Query("roleType")
	roleType, _ := strconv.Atoi(roleTypeS)
	registerCode := ctx.Query("registerCode")
	deptIdS := ctx.Query("deptId")
	deptId, _ := strconv.Atoi(deptIdS)
	phoneNumber := ctx.Query("phoneNumber")
	flag := user_service.RegisterUser(username, password, registerCode, phoneNumber, roleType, deptId)
	if flag {
		response.Success(ctx, "注册成功！")
		return
	}
	response.Fail(ctx, "注册码有误！")
}

func DeptList(ctx *gin.Context) {
	dept, _ := service.GetAllDept("true")
	response.Success(ctx, dept)
}
