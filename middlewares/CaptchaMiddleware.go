package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
)

func CaptchaMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		validate(context)
	}
}

func validate(ctx *gin.Context) {
	code := ctx.Query("code")
	key := ctx.Query("key")
	if code == "" || key == "" {
		panic(errors.New("验证码错误"))
	}
	if flag := common.VerifyCaptcha(key, code); !flag {
		panic(errors.New("验证码错误"))
	}
}
