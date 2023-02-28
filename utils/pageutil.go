package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(ctx *gin.Context) (offset int, limit int) {
	pageNum := 1
	pageSize := 10
	pnStr := ctx.Param("pageNum")
	if pnStr != "" {
		p, err := strconv.Atoi(pnStr)
		if err == nil {
			pageNum = p
		}
	}
	psStr := ctx.Param("pageSize")
	if psStr != "" {
		p, err := strconv.Atoi(psStr)
		if err == nil {
			pageSize = p
		}
	}
	return (pageNum - 1) * pageSize, pageSize
}
