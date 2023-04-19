package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
)

func GetListHealthClock(ctx *gin.Context) {
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
	healthClockQ := query.Use(mysql.DB).HealthClock
	if flag {
		healthTypeS := ctx.Query("health_type")
		if healthTypeS != "" {
			atoi, _ := strconv.Atoi(healthTypeS)
			cds = append(cds, healthClockQ.HealthType.Eq(int32(atoi)))
		}
		middle_highS := ctx.Query("middle_high")
		if middle_highS != "" {
			atoi, _ := strconv.Atoi(middle_highS)
			cds = append(cds, healthClockQ.MiddleHigh.Eq(int32(atoi)))
		}
		diagnosisS := ctx.Query("diagnosis")
		if diagnosisS != "" {
			atoi, _ := strconv.Atoi(diagnosisS)
			cds = append(cds, healthClockQ.Diagnosis.Eq(int32(atoi)))
		}
		returnInfoS := ctx.Query("return_info")
		if returnInfoS != "" {
			atoi, _ := strconv.Atoi(returnInfoS)
			cds = append(cds, healthClockQ.ReturnInfo.Eq(int32(atoi)))
		}
		deptIdS := ctx.Query("deptId")
		if deptIdS != "" {
			atoi, _ := strconv.Atoi(deptIdS)
			cds = append(cds, healthClockQ.DeptID.Eq(int32(atoi)))
		}
	} else {
		cds = append(cds, healthClockQ.Username.Eq(username))
	}
	offset, limit := utils.GetPage(ctx)
	data, count, err := healthClockQ.WithContext(context.Background()).Where(cds...).Order(healthClockQ.CreateTime.Desc()).FindByPage(offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func CheckHealthClock(ctx *gin.Context) {
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	count := service.CheckClockToday(username)
	if count > 0 {
		response.Success(ctx, "今日已打卡，请不要重复打卡！")
		return
	} else {
		response.Success(ctx, "allow")
		return
	}
}

func SaveHealthClock(ctx *gin.Context) {
	healthClock := new(models.HealthClock)
	err := ctx.ShouldBindJSON(healthClock)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	healthClockQ := query.Use(mysql.DB).HealthClock
	err = healthClockQ.WithContext(context.Background()).Create(healthClock)
	if err != nil {
		response.Fail(ctx, "打卡失败")
		return
	}
	response.Success(ctx, "打卡成功")
}
