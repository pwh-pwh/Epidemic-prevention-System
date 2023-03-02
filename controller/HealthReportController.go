package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
)

func CheckHealthReport(ctx *gin.Context) {
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	count := service.CheckReportToday(username)
	if count > 0 {
		response.Success(ctx, "今日已上报，请不要重复上报！")
		return
	} else {
		response.Success(ctx, "allow")
	}
}

func ListHealthReport(ctx *gin.Context) {
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	//判断是否是教师
	auListI, _ := ctx.Get("auList")
	authorityList := auListI.([]string)
	flag := false
	for _, item := range authorityList {
		if item == common.RoleTeacher {
			flag = true
		}
	}
	var cds []gen.Condition
	healthReportQ := query.Use(mysql.DB).HealthReport
	if flag {
		name := ctx.Query("username")
		if name != "" {
			cds = append(cds, healthReportQ.Username.Like("%"+name+"%"))
		}
		deptIdS := ctx.Query("deptId")
		if deptIdS != "" {
			atoi, _ := strconv.Atoi(deptIdS)
			cds = append(cds, healthReportQ.DeptID.Eq(int64(atoi)))
		}
		typeS := ctx.Query("type")
		if typeS != "" {
			atoi, _ := strconv.Atoi(typeS)
			cds = append(cds, healthReportQ.Type.Eq(int32(atoi)))
		}
		start := ctx.Query("start")
		end := ctx.Query("end")
		if start != "" && end != "" {
			st, _ := utils.ParseTime(start+" 00:00:00", common.TimeFormat)
			et, _ := utils.ParseTime(end+" 23:59:59", common.TimeFormat)
			cds = append(cds, healthReportQ.CreateTime.Between(st, et))
		}
	} else {
		cds = append(cds, healthReportQ.Username.Eq(username))
	}
	offset, limit := utils.GetPage(ctx)
	data, count, err := healthReportQ.WithContext(context.Background()).Where(cds...).Order(healthReportQ.CreateTime.Desc()).FindByPage(offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func SaveHealthReport(ctx *gin.Context) {
	file1, _ := ctx.FormFile("file1")
	file2, _ := ctx.FormFile("file2")
	file3, _ := ctx.FormFile("file3")
	healthReport := new(models.HealthReport)
	if file3.Size != 0 && file2.Size != 0 && file1.Size != 0 {
		img1, _ := utils.UploadFile(file1)
		img2, _ := utils.UploadFile(file2)
		img3, _ := utils.UploadFile(file3)
		healthReport.Img1 = img1
		healthReport.Img2 = img2
		healthReport.Img3 = img3
	}
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	redisClient := redis.GetRedisClient()
	sysUser := new(models.SysUser)
	err := redisClient.Get(redis.UserPre + username).Scan(sysUser)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	healthReport.Username = username
	healthReport.PhoneNumber = sysUser.PhoneNumber
	healthReport.DeptID = sysUser.DeptID
	err = query.Use(mysql.DB).HealthReport.WithContext(context.Background()).Create(healthReport)
	if err != nil {
		response.Fail(ctx, "上传失败")
		return
	}
	response.Success(ctx, "上报成功")
}
