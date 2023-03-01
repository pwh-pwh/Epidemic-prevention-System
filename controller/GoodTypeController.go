package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
	"strings"
)

func GetSimpleListGoodsType(ctx *gin.Context) {
	goodsTypeQ := query.Use(mysql.DB).GoodType
	var data []struct {
		Id   int64  `json:"id"`
		Type string `json:"type"`
	}
	err := goodsTypeQ.WithContext(context.Background()).Select(goodsTypeQ.ID, goodsTypeQ.Type).
		Where(goodsTypeQ.Status.Eq(1)).
		Order(goodsTypeQ.OrderNum).Scan(&data)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, data)
}

func GetListGoodsType(ctx *gin.Context) {
	goodsTypeQ := query.Use(mysql.DB).GoodType
	typeStr := ctx.Query("type")
	var cds []gen.Condition
	if typeStr != "" {
		cds = append(cds, goodsTypeQ.Type.Like("%"+typeStr+"%"))
	}
	createBy := ctx.Query("createBy")
	if createBy != "" {
		cds = append(cds, goodsTypeQ.CreateBy.Like("%"+createBy+"%"))
	}
	statusStr := ctx.Query("status")
	if statusStr != "" {
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			response.Fail(ctx, err.Error())
			return
		}
		cds = append(cds, goodsTypeQ.Status.Eq(int32(status)))
	}
	data, count, err := goodsTypeQ.WithContext(context.Background()).Where(cds...).Order(goodsTypeQ.OrderNum).FindByPage(utils.GetPage(ctx))
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func SaveGoodsType(ctx *gin.Context) {
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	goodsType := new(models.GoodType)
	err := ctx.ShouldBindJSON(goodsType)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	goodsType.CreateBy = username
	goodsTypeQ := query.Use(mysql.DB).GoodType
	err = goodsTypeQ.WithContext(context.Background()).Create(goodsType)
	if err != nil {
		response.Fail(ctx, "添加类型失败！")
		return
	}
	response.Success(ctx, "添加类型成功！")
}

func UpdateGoodsType(ctx *gin.Context) {
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	goodsType := new(models.GoodType)
	err := ctx.ShouldBindJSON(goodsType)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	goodsType.CreateBy = username
	goodsTypeQ := query.Use(mysql.DB).GoodType
	_, err = goodsTypeQ.WithContext(context.Background()).Select(goodsTypeQ.Type, goodsTypeQ.Status, goodsTypeQ.CreateBy, goodsTypeQ.Remark, goodsTypeQ.OrderNum).Where(goodsTypeQ.ID.Eq(goodsType.ID)).Updates(goodsType)
	if err != nil {
		response.Fail(ctx, "修改类型失败！")
		return
	}
	response.Success(ctx, "修改类型成功！")
}

func DeleteGoodsType(ctx *gin.Context) {
	idsS := ctx.Query("ids")
	split := strings.Split(idsS, ",")
	var ids []int64
	for _, s := range split {
		parseInt, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, parseInt)
	}
	goodsTypeQ := query.Use(mysql.DB).GoodType
	_, err := goodsTypeQ.WithContext(context.Background()).Where(goodsTypeQ.ID.In(ids...)).Delete()
	if err != nil {
		response.Fail(ctx, "删除类型失败！")
		return
	}
	response.Success(ctx, "删除类型成功！")
}
