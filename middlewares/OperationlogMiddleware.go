package middlewares

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/task"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"io/ioutil"
	"net/http"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyLogWriter) Write(data []byte) (int, error) {
	w.bodyBuf.Write(data)
	return w.ResponseWriter.Write(data)
}

func CommonLogInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logT(ctx)
	}
}

func logT(ctx *gin.Context) {
	method := ctx.Request.Method
	if method == http.MethodGet || method == http.MethodOptions {
		return
	}
	if ctx.Request.URL.Path == "/login" {
		return
	}
	opLog := new(models.SysOperateLog)
	opLog.RequestMethod = ctx.Request.Method
	opLog.OperURL = ctx.Request.URL.Path
	opLog.OperIP = ctx.ClientIP()
	opLog.OperLocation = utils.GetLocation(ctx.ClientIP())
	opLog.OperTime = models.LocalTime(time.Now())

	var bodyBytes []byte
	if ctx.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	opLog.OperParam = string(bodyBytes)
	var blw bodyLogWriter
	blw = bodyLogWriter{bodyBuf: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw
	ctx.Next()
	strBody := blw.bodyBuf.String()
	//文件上传忽略
	if len(strBody) > 1024*1024 {
		return
	}
	opLog.OperParam = strBody
	usernameI, ise := ctx.Get("username")
	if ise {
		opLog.OperName = usernameI.(string)
	}
	typeI, exists := ctx.Get("type")
	if exists {
		opLog.BusinessType = typeI.(string)
	}
	titleI, exists := ctx.Get("title")
	if exists {
		opLog.Title = titleI.(string)
	}
	me, exists := ctx.Get("method")
	if exists {
		opLog.Method = me.(string)
	}
	defer func() {
		var err error
		if e := recover(); e != nil {
			err = e.(error)
			opLog.ErrorMsg = err.Error()
			//panic(err)
		}
		task.AddTask(func() {
			logQ := query.Use(mysql.DB).SysOperateLog
			_ = logQ.WithContext(context.Background()).Create(opLog)
		})
		if err != nil {
			panic(err)
		}
	}()
}
