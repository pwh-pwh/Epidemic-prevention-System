package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
)

func GetCaptcha(c *gin.Context) {
	img, id := common.NewBase64Img()
	response.Success(c, gin.H{
		"base64": img,
		"key":    id,
	})
}
