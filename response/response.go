package response

import "github.com/gin-gonic/gin"

type response struct {
	code int
	msg  string
	data any
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(200, response{
		code: 200,
		msg:  "操作成功",
		data: data,
	})
}

func Fail(ctx *gin.Context, msg string) {
	ctx.JSON(200, response{
		code: 400,
		msg:  msg,
		data: nil,
	})
}

func Response(ctx *gin.Context, code int, msg string, data any) {
	ctx.JSON(200, response{
		code: code,
		msg:  msg,
		data: data,
	})
}
