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
)

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

/*
@GetMapping("/list")
    @PreAuthorize("hasAnyAuthority('good:stock:list')")
    public Result list(String accept, Integer operateType, String start, String end) {
        LambdaQueryWrapper<GoodStock> wrapper = Wrappers.lambdaQuery(GoodStock.class);
        wrapper.like(StrUtil.isNotBlank(accept), GoodStock::getAccept, accept);
        wrapper.eq(operateType != null, GoodStock::getOperateType, operateType);
        if (StrUtil.isNotBlank(start) && StrUtil.isNotBlank(end)){
            DateTime a = DateUtil.parse(start + " 00:00:00", "yyyy-MM-dd HH:mm:ss");
            DateTime b = DateUtil.parse(end + " 23:59:59", "yyyy-MM-dd HH:mm:ss");
            wrapper.between(GoodStock::getCreateTime, a, b);
        }
        wrapper.orderByDesc(GoodStock::getCreateTime);
        Page<GoodStock> page = goodStockService.page(getPage(), wrapper);
        return Result.succ(page);
    }

    @PostMapping
    @Log(title = "物资库存", businessType = "物资出入库")
    @PreAuthorize("hasAnyAuthority('good:stock:operate')")
    public Result save(@Validated @RequestBody GoodStock goodStock, Principal principal) {
        List<GoodDto> list = goodStock.getList();
        List<GoodStock> goodStockList = new ArrayList<>();
        list.forEach(goodDto -> {
            GoodStock stock = new GoodStock();
            stock.setCreateBy(principal.getName());
            stock.setAccept(goodStock.getAccept());
            stock.setGoodName(goodDto.getGoodName());
            stock.setGoodNum(goodDto.getGoodNum());
            stock.setGoodSize(goodDto.getGoodSize());
            stock.setPeopleName(goodStock.getPeopleName());
            stock.setPeoplePhone(goodStock.getPeoplePhone());
            stock.setOperateType(goodStock.getOperateType());
            stock.setRemark(goodStock.getRemark());
            goodStockList.add(stock);
            GoodInfo goodInfo = goodInfoService.getById(goodDto.getId());
            int res;
            if (goodStock.getOperateType() == 0) {
                res = goodInfo.getTotal() + goodDto.getGoodNum();
            } else {
                res = goodInfo.getTotal() - goodDto.getGoodNum();
            }
            goodInfo.setTotal(res);
            goodInfoService.updateById(goodInfo);
        });
        boolean batch = goodStockService.saveBatch(goodStockList);
        if (batch) {
            return Result.succ("操作成功！");
        } else {
            return Result.fail("操作失败！");
        }
    }
*/
