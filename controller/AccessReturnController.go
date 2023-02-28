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

/*
@GetMapping("/list")
    @PreAuthorize("hasAnyAuthority('access:return:list')")
    public Result list(String name, String dept, String start, String end){
        LambdaQueryWrapper<AccessReturn> wrapper = Wrappers.lambdaQuery(AccessReturn.class);
        wrapper.like(StrUtil.isNotBlank(name), AccessReturn::getName, name);
        wrapper.like(StrUtil.isNotBlank(dept), AccessReturn::getDept, dept);
        if (StrUtil.isNotBlank(start) && StrUtil.isNotBlank(end)){
            DateTime a = DateUtil.parse(start + " 00:00:00", "yyyy-MM-dd HH:mm:ss");
            DateTime b = DateUtil.parse(end + " 23:59:59", "yyyy-MM-dd HH:mm:ss");
            wrapper.between(AccessReturn::getCreateTime, a, b);
        }
        wrapper.orderByDesc(AccessReturn::getCreateTime);
        Page<AccessReturn> page = accessReturnService.page(getPage(), wrapper);
        return Result.succ(page);
    }
*/

func AccessReturnList(ctx *gin.Context) {
	name := ctx.Param("name")
	dept := ctx.Param("dept")
	start := ctx.Param("start")
	end := ctx.Param("end")
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
