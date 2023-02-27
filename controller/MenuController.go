package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/menu_service"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/user_service"
)

func Nav(ctx *gin.Context) {
	auListI, ok := ctx.Get("auList")
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	var auList []string
	if ok {
		auList = auListI.([]string)
	} else {
		auList = user_service.GetUserAuthorityList(username)
	}
	//获取nav
	nav, err := menu_service.GetUserNav(username)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{
		"authoritys": auList,
		"nav":        nav,
	})
}
