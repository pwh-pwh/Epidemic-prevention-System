package response

import "github.com/gin-gonic/gin"

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(200, response{
		Code: 200,
		Msg:  "操作成功",
		Data: data,
	})
}

func Fail(ctx *gin.Context, msg string) {
	ctx.JSON(200, response{
		Code: 400,
		Msg:  msg,
		Data: nil,
	})
}

func Response(ctx *gin.Context, code int, msg string, data any) {
	ctx.JSON(200, response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
