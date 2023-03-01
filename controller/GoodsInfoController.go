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

//TODO
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

/*
   @PutMapping
   @Log(title = "物资资料", businessType = "修改物资")
   @PreAuthorize("hasAnyAuthority('good:info:update')")
   public Result update(@RequestParam(value = "img", required = false) MultipartFile file, @RequestParam("goodInfo") String json) throws JsonProcessingException {
       ObjectMapper mapper = new ObjectMapper();
       GoodInfo goodInfo = mapper.readValue(json, GoodInfo.class);
       if (file != null && !file.isEmpty()) {
           String uploadImg = UploadUtil.uploadImg(file);
           if (StrUtil.isEmpty(uploadImg)) {
               return Result.fail("图片上传失败");
           }
           goodInfo.setImg(Const.IMG_PATH + uploadImg);
       }
       boolean update = goodInfoService.updateById(goodInfo);
       if (update) {
           return Result.succ("修改成功");
       } else {
           return Result.fail("修改失败");
       }
   }

   @DeleteMapping
   @Log(title = "物资资料", businessType = "删除物资")
   @PreAuthorize("hasAnyAuthority('good:info:delete')")
   public Result delete(Long[] ids) {
       boolean remove = goodInfoService.removeByIds(Arrays.asList(ids));
       if (remove) {
           return Result.succ("删除成功！");
       } else {
           return Result.fail("删除成功！");
       }
   }

*/
