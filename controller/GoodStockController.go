package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
)

// GetListGoodsStock GetListGoodsStock接口
// @Summary GetListGoodsStock接口
// @Description 可按accept按operateType按dept或start,end,根据createTime排序查询列表接口
// @Tags goodStock相关接口
// @Produce application/json
// @Param Authorization header string false "jwt"
// @Security ApiKeyAuth
// @Success 200 {object} response.response
// @Router /good/stock/list [get]
func GetListGoodsStock(ctx *gin.Context) {
	goodsStockQ := query.Use(mysql.DB).GoodStock
	var cds []gen.Condition
	accept := ctx.Query("accept")
	if accept != "" {
		cds = append(cds, goodsStockQ.Accept.Like("%"+accept+"%"))
	}
	operateTypeS := ctx.Query("operateType")
	if operateTypeS != "" {
		atoi, err := strconv.Atoi(operateTypeS)
		if err != nil {
			response.Fail(ctx, err.Error())
			return
		}
		cds = append(cds, goodsStockQ.OperateType.Eq(int32(atoi)))
	}
	start := ctx.Query("start")
	end := ctx.Query("end")
	if start != "" && end != "" {
		startTime, _ := utils.ParseTime(start+" 00:00:00", common.TimeFormat)
		endTime, _ := utils.ParseTime(end+" 00:00:00", common.TimeFormat)
		cds = append(cds, goodsStockQ.CreateTime.Between(startTime, endTime))
	}
	offset, limit := utils.GetPage(ctx)
	data, count, err := goodsStockQ.WithContext(context.Background()).Where(cds...).Order(goodsStockQ.CreateTime.Desc()).FindByPage(offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func SaveGoodsStock(ctx *gin.Context) {
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	goodsStock := new(models.GoodStock)
	goodStockList := make([]*models.GoodStock, 0)
	err := ctx.ShouldBindJSON(goodsStock)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	goodsInfoQ := query.Use(mysql.DB).GoodInfo
	for _, item := range goodsStock.List {
		st := new(models.GoodStock)
		st.CreateBy = username
		st.Accept = goodsStock.Accept
		st.GoodName = item.GoodName
		st.GoodNum = item.GoodNum
		st.GoodSize = item.GoodSize
		st.PeopleName = goodsStock.PeopleName
		st.PeoplePhone = goodsStock.PeoplePhone
		st.OperateType = goodsStock.OperateType
		st.Remark = goodsStock.Remark
		goodStockList = append(goodStockList, st)
		var res int32
		goodsInfo, _ := goodsInfoQ.WithContext(context.Background()).Where(goodsInfoQ.ID.Eq(item.Id)).Take()
		if goodsStock.OperateType == 0 {
			res = goodsInfo.Total + item.GoodNum
		} else {
			res = goodsInfo.Total - item.GoodNum
		}
		goodsInfo.Total = res
		_, _ = goodsInfoQ.WithContext(context.Background()).Where(goodsInfoQ.ID.Eq(goodsInfo.ID)).Updates(goodsInfo)
	}
	goodsStockQ := query.Use(mysql.DB).GoodStock
	err = goodsStockQ.WithContext(context.Background()).Create(goodStockList...)
	if err != nil {
		response.Fail(ctx, "操作失败！")
		return
	}
	response.Success(ctx, "操作成功！")
}
