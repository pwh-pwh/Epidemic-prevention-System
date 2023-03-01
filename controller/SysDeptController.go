package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service"
)

func ListDept(ctx *gin.Context) {
	flagS := ctx.Param("flag")
	depts, err := service.GetAllDept(flagS)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, depts)
}
