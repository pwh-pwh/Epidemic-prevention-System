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

/*
   @GetMapping("/total")
   @PreAuthorize("hasAnyAuthority('good:info:list')")
   public Result total(String name, Long typeId) {
       QueryWrapper<GoodInfo> wrapper = new QueryWrapper<>();
       wrapper.like(StrUtil.isNotBlank(name), "name", name);
       wrapper.eq(typeId != null, "type_id", typeId);
       wrapper.select("id", "name", "unit", "size", "total");
       Page<GoodInfo> page = goodInfoService.page(getPage(), wrapper);
       return Result.succ(page);
   }

   @GetMapping("/list")
   @PreAuthorize("hasAnyAuthority('good:info:list')")
   public Result list(String name, String createBy, Integer status, Long typeId) {
       QueryWrapper<GoodInfo> wrapper = new QueryWrapper<>();
       wrapper.like(StrUtil.isNotBlank(name), "name", name);
       wrapper.like(StrUtil.isNotBlank(createBy), "create_by", createBy);
       wrapper.eq(status != null, "status", status);
       wrapper.eq(typeId != null, "type_id", typeId);
       Page<GoodInfo> page = goodInfoService.page(getPage(), wrapper);
       return Result.succ(page);
   }
*/
