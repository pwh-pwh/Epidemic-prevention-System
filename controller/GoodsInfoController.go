package controller

import (
	"context"
	"encoding/json"
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

func SaveGoodsInfo(ctx *gin.Context) {
	imgFile, err := ctx.FormFile("img")
	if err != nil {
		response.Fail(ctx, "请上传图片")
		return
	}
	goodsInfoJsonStr := ctx.PostForm("goodInfo")
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	goodsInfo := new(models.GoodInfo)
	err = json.Unmarshal([]byte(goodsInfoJsonStr), goodsInfo)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	imgUrl, err := utils.UploadFile(imgFile)
	if err != nil {
		response.Fail(ctx, "图片上传失败")
		return
	}
	goodsInfo.Img = imgUrl
	goodsInfo.CreateBy = username
	goodsInfoQ := query.Use(mysql.DB).GoodInfo
	err = goodsInfoQ.WithContext(context.Background()).Create(goodsInfo)
	if err != nil {
		response.Fail(ctx, "新增失败")
		return
	}
	response.Success(ctx, "新增成功")
}

func UpdateGoodsInfo(ctx *gin.Context) {
	goodsInfoJsonStr := ctx.PostForm("goodInfo")
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	goodsInfo := new(models.GoodInfo)
	err := json.Unmarshal([]byte(goodsInfoJsonStr), goodsInfo)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	imgFile, err := ctx.FormFile("img")
	if err == nil {
		imgUrl, err := utils.UploadFile(imgFile)
		if err != nil {
			response.Fail(ctx, "图片上传失败")
			return
		}
		goodsInfo.Img = imgUrl
	}
	goodsInfo.CreateBy = username
	goodsInfoQ := query.Use(mysql.DB).GoodInfo
	verCd := goodsInfoQ.Version.Eq(goodsInfo.Version)
	goodsInfo.Version += 1
	_, err = goodsInfoQ.WithContext(context.Background()).Where(goodsInfoQ.ID.Eq(goodsInfo.ID), verCd).Updates(goodsInfo)
	if err != nil {
		response.Fail(ctx, "修改失败")
		return
	}
	response.Success(ctx, "修改成功")
}

func DeleteGoodsInfo(ctx *gin.Context) {
	var ids []int64
	idsStr := ctx.Query("ids")
	split := strings.Split(idsStr, ",")
	for _, s := range split {
		atoi, _ := strconv.Atoi(s)
		ids = append(ids, int64(atoi))
	}
	goodsInfoQ := query.Use(mysql.DB).GoodInfo
	_, err := goodsInfoQ.WithContext(context.Background()).Where(goodsInfoQ.ID.In(ids...)).Delete()
	if err != nil {
		response.Fail(ctx, "删除失败！")
		return
	}
	response.Success(ctx, "删除成功！")
}
