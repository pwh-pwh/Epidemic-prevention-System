package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
)

// GetAccessRegisterList GetAccessRegisterList接口
// @Summary GetAccessRegisterList接口
// @Description 可按type按name或start,end,根据createTime排序查询列表接口
// @Tags ccessRegister相关接口
// @Produce application/json
// @Param Authorization header string false "jwt"
// @Security ApiKeyAuth
// @Success 200 {object} response.response
// @Router /access/register/list [get]
func GetAccessRegisterList(ctx *gin.Context) {
	name := ctx.Query("name")
	typeStr := ctx.Query("type")
	start := ctx.Query("start")
	end := ctx.Query("end")
	var cd []gen.Condition
	accessRegister := query.Use(mysql.DB).AccessRegister
	if typeStr != "" {
		typeInt, _ := strconv.Atoi(typeStr)
		cd = append(cd, accessRegister.Type.Eq(int32(typeInt)))
	}
	if name != "" {
		cd = append(cd, accessRegister.Name.Like("%"+name+"%"))
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
		cd = append(cd, accessRegister.CreateTime.Between(sT, eT))
	}
	offset, limit := utils.GetPage(ctx)
	result, count, err := accessRegister.WithContext(context.Background()).Where(cd...).Order(accessRegister.CreateTime.Desc()).FindByPage(offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: result,
		Total:   count,
	})
}

func SaveAccessRegister(ctx *gin.Context) {
	ar := new(models.AccessRegister)
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	ar.CreateBy = username
	if err := ctx.ShouldBindJSON(ar); err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	flag := service.AddAccessRegister(ar)
	if flag {
		response.Success(ctx, "登记成功")
	} else {
		response.Fail(ctx, "登记失败")
	}
}
