package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
)

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.Fail(ctx, "请上传文件")
		return
	}
	path, err := utils.UploadFile(file)
	if err != nil {
		response.Fail(ctx, "图片上传失败")
		return
	}
	response.Success(ctx, path)
}
