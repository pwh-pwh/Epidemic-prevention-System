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
)

// AccessReturnList AccessReturnList接口
// @Summary AccessReturnList接口
// @Description 可按type按name按dept或start,end,根据createTime排序查询列表接口
// @Tags ccessReturn相关接口
// @Produce application/json
// @Param Authorization header string false "jwt"
// @Security ApiKeyAuth
// @Success 200 {object} response.response
// @Router /access/return/list [get]
func AccessReturnList(ctx *gin.Context) {
	name := ctx.Query("name")
	dept := ctx.Query("dept")
	start := ctx.Query("start")
	end := ctx.Query("end")
	var cd []gen.Condition

	accessReturn := query.Use(mysql.DB).AccessReturn
	if name != "" {
		cd = append(cd, accessReturn.Name.Like("%"+name+"%"))
	}
	if dept != "" {
		cd = append(cd, accessReturn.Dept.Like("%"+dept+"%"))
	}
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
		cd = append(cd, accessReturn.CreateTime.Between(sT, eT))
	}
	offset, limit := utils.GetPage(ctx)
	result, count, err := accessReturn.WithContext(context.Background()).Where(cd...).Order(accessReturn.CreateTime).FindByPage(offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: result,
		Total:   count,
	})
}
