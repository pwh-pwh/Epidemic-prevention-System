package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
)

func GetAllGoodsInfo(ctx *gin.Context) {
	goodInfoQ := query.Use(mysql.DB).GoodInfo
	var data []vo.GoodsInfoVO
	err := goodInfoQ.WithContext(context.Background()).Select(goodInfoQ.ID, goodInfoQ.Name, goodInfoQ.Unit,
		goodInfoQ.Size, goodInfoQ.Total).Scan(&data)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, data)
}

/*
 @GetMapping("/all")
    @PreAuthorize("hasAnyAuthority('good:info:list')")
    public Result all(){
        QueryWrapper<GoodInfo> wrapper = new QueryWrapper<>();
        wrapper.select("id", "name", "unit", "size", "total");
        List<GoodInfo> list = goodInfoService.list(wrapper);
        return Result.succ(list);
    }

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
