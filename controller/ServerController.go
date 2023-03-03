package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dto"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
)

func ServerInfo(ctx *gin.Context) {
	s := dto.GetServer()
	response.Success(ctx, s)
}
