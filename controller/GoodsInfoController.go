package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
)

func GetAllGoodsInfo(ctx *gin.Context) {
	goodsInfoQ := query.Use(mysql.DB).GoodInfo
	var data []vo.GoodsInfoVO
	err := goodsInfoQ.WithContext(context.Background()).Select(goodsInfoQ.ID, goodsInfoQ.Name, goodsInfoQ.Unit,
		goodsInfoQ.Size, goodsInfoQ.Total).Scan(&data)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, data)
}

func GetTotalGoodsInfo(ctx *gin.Context) {
	goodsInfoQ := query.Use(mysql.DB).GoodInfo
	var cds []gen.Condition
	name := ctx.Query("name")
	if name != "" {
		cds = append(cds, goodsInfoQ.Name.Like("%"+name+"%"))
	}
	typeIdS := ctx.Query("typeId")
	if typeIdS != "" {
		typeId, err := strconv.Atoi(typeIdS)
		if err != nil {
			response.Fail(ctx, err.Error())
			return
		}
		cds = append(cds, goodsInfoQ.TypeID.Eq(int64(typeId)))
	}
	offset, limit := utils.GetPage(ctx)
	var data []vo.GoodsInfoVO
	count, err := goodsInfoQ.WithContext(context.Background()).Select(goodsInfoQ.ID, goodsInfoQ.Name, goodsInfoQ.Unit,
		goodsInfoQ.Size, goodsInfoQ.Total).Where(cds...).ScanByPage(&data, offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func GetListGoodsInfo(ctx *gin.Context) {
	goodsInfoQ := query.Use(mysql.DB).GoodInfo
	var cds []gen.Condition
	name := ctx.Query("name")
	if name != "" {
		cds = append(cds, goodsInfoQ.Name.Like("%"+name+"%"))
	}
	typeIdS := ctx.Query("typeId")
	if typeIdS != "" {
		typeId, err := strconv.Atoi(typeIdS)
		if err != nil {
			response.Fail(ctx, err.Error())
			return
		}
		cds = append(cds, goodsInfoQ.TypeID.Eq(int64(typeId)))
	}
	createBy := ctx.Query("createBy")
	if createBy != "" {
		cds = append(cds, goodsInfoQ.CreateBy.Like("%"+createBy+"%"))
	}
	statusS := ctx.Query("status")
	if statusS != "" {
		status, err := strconv.Atoi(statusS)
		if err != nil {
			response.Fail(ctx, err.Error())
			return
		}
		cds = append(cds, goodsInfoQ.TypeID.Eq(int64(status)))
	}
	offset, limit := utils.GetPage(ctx)
	data, count, err := goodsInfoQ.WithContext(context.Background()).Where(cds...).FindByPage(offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}
