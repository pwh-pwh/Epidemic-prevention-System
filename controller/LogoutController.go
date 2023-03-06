package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
)

func Logout(ctx *gin.Context) {
	ctx.Header(common.GetJwtHeader(), "")
	response.Success(ctx, "退出成功")
}
