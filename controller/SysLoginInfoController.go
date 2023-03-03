package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
	"strings"
)

func ListLoginInfo(ctx *gin.Context) {
	var cds []gen.Condition
	loginInfoQ := query.Use(mysql.DB).SysLoginInfo
	ip := ctx.Query("ip")
	if ip != "" {
		cds = append(cds, loginInfoQ.IP.Like("%"+ip+"%"))
	}
	username := ctx.Query("username")
	if username != "" {
		cds = append(cds, loginInfoQ.Username.Like("%"+username+"%"))
	}
	statusS := ctx.Query("status")
	if statusS != "" {
		status, _ := strconv.Atoi(statusS)
		cds = append(cds, loginInfoQ.Status.Eq(int32(status)))
	}
	start := ctx.Query("start")
	end := ctx.Query("end")
	if start != "" && end != "" {
		sT, err := utils.ParseTime(start+" 00:00:00", common.TimeFormat)
		if err != nil {
			response.Fail(ctx, "时间参数错误")
			return
		}
		eT, err := utils.ParseTime(end+" 23:59:59", common.TimeFormat)
		if err != nil {
			response.Fail(ctx, "时间参数错误")
			return
		}
		cds = append(cds, loginInfoQ.LoginTime.Between(sT, eT))
	}
	data, count, err := loginInfoQ.WithContext(context.Background()).Where(cds...).Order(loginInfoQ.LoginTime.Desc()).FindByPage(utils.GetPage(ctx))
	if err != nil {
		response.Fail(ctx, "获取列表失败")
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func DeleteLoginInfo(ctx *gin.Context) {
	idsS := ctx.Query("ids")
	split := strings.Split(idsS, ",")
	var ids []int64
	for _, s := range split {
		parseInt, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, parseInt)
	}
	sysLoginQ := query.Use(mysql.DB).SysLoginInfo
	_, err := sysLoginQ.WithContext(context.Background()).Where(sysLoginQ.ID.In(ids...)).Delete()
	if err != nil {
		response.Fail(ctx, "删除失败！")
		return
	}
	response.Success(ctx, "删除成功！")
}

func ClearLoginInfo(ctx *gin.Context) {
	mysql.DB.Exec("truncate table sys_login_info")
	response.Success(ctx, "所有数据已清除")
}
