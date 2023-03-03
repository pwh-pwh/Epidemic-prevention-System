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

func ListOpLog(ctx *gin.Context) {
	var cds []gen.Condition
	opLogQ := query.Use(mysql.DB).SysOperateLog
	title := ctx.Query("title")
	if title != "" {
		cds = append(cds, opLogQ.Title.Like("%"+title+"%"))
	}
	operName := ctx.Query("operName")
	if operName != "" {
		cds = append(cds, opLogQ.OperName.Like("%"+operName+"%"))
	}
	statusS := ctx.Query("status")
	if statusS != "" {
		status, _ := strconv.Atoi(statusS)
		cds = append(cds, opLogQ.Status.Eq(int32(status)))
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
		cds = append(cds, opLogQ.OperTime.Between(sT, eT))
	}
	offset, limit := utils.GetPage(ctx)
	data, count, err := opLogQ.WithContext(context.Background()).Where(cds...).Order(opLogQ.OperTime.Desc()).FindByPage(offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func DeleteOpLog(ctx *gin.Context) {
	idsS := ctx.Query("ids")
	split := strings.Split(idsS, ",")
	var ids []int64
	for _, s := range split {
		parseInt, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, parseInt)
	}
	opLogQ := query.Use(mysql.DB).SysOperateLog
	_, err := opLogQ.WithContext(context.Background()).Where(opLogQ.ID.In(ids...)).Delete()
	if err != nil {
		response.Fail(ctx, "删除失败")
		return
	}
	response.Success(ctx, "删除成功")
}

func ClearOpLog(ctx *gin.Context) {
	mysql.DB.Exec("truncate table sys_operate_log")
	response.Success(ctx, "清空成功")
}
