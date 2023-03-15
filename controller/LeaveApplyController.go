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
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
)

// GetListLeaveApply GetListLeaveApply接口
// @Summary GetListLeaveApply接口
// @Description 可按status按username按deptId或start,end,根据createTime排序查询列表接口
// @Tags leaveApply相关接口
// @Produce application/json
// @Param Authorization header string false "jwt"
// @Security ApiKeyAuth
// @Success 200 {object} response.response
// @Router /leave/apply/list [get]
func GetListLeaveApply(ctx *gin.Context) {
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
	leaveApplyQ := query.Use(mysql.DB).LeaveApply
	if flag {
		username := ctx.Query("username")
		if username != "" {
			cds = append(cds, leaveApplyQ.Username.Like("%"+username+"%"))
		}
		deptIdS := ctx.Query("deptId")
		if deptIdS != "" {
			atoi, _ := strconv.Atoi(deptIdS)
			cds = append(cds, leaveApplyQ.DeptID.Eq(int64(atoi)))
		}
		statusS := ctx.Query("status")
		if statusS != "" {
			atoi, _ := strconv.Atoi(statusS)
			cds = append(cds, leaveApplyQ.Status.Eq(int32(atoi)))
		}
		start := ctx.Query("start")
		end := ctx.Query("end")
		if start != "" && end != "" {
			st, _ := utils.ParseTime(start+" 00:00:00", common.TimeFormat)
			et, _ := utils.ParseTime(end+" 23:59:59", common.TimeFormat)
			cds = append(cds, leaveApplyQ.CreateTime.Between(st, et))
		}
	} else {
		cds = append(cds, leaveApplyQ.Username.Eq(username))
	}
	offset, limit := utils.GetPage(ctx)
	data, count, err := leaveApplyQ.WithContext(context.Background()).Where(cds...).Order(leaveApplyQ.CreateTime.Desc()).FindByPage(offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func SaveLeaveApply(ctx *gin.Context) {
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	redisClient := redis.GetRedisClient()
	sysUser := new(models.SysUser)
	_ = redisClient.Get(redis.UserPre + username).Scan(sysUser)
	leaveApply := new(models.LeaveApply)
	err := ctx.ShouldBindJSON(leaveApply)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	leaveApply.Status = 1
	leaveApply.Username = username
	leaveApply.DeptID = sysUser.DeptID
	leaveApply.Nickname = sysUser.Nickname
	if err = query.Use(mysql.DB).LeaveApply.WithContext(context.Background()).Create(leaveApply); err != nil {
		response.Fail(ctx, "请假提交失败！")
		return
	}
	response.Success(ctx, "请假提交成功！")
}

func UpdateLeaveApply(ctx *gin.Context) {
	leaveApply := new(models.LeaveApply)
	err := ctx.ShouldBindJSON(leaveApply)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	leaveApplyQ := query.Use(mysql.DB).LeaveApply
	_, err = leaveApplyQ.WithContext(context.Background()).
		Select(leaveApplyQ.LeaveType, leaveApplyQ.StudentType, leaveApplyQ.Time, leaveApplyQ.Status,
			leaveApplyQ.Day, leaveApplyQ.Address, leaveApplyQ.Traffic, leaveApplyQ.Dormitory, leaveApplyQ.PhoneNumber, leaveApplyQ.Clazz, leaveApplyQ.Exam, leaveApplyQ.Reason).
		Where(leaveApplyQ.ID.Eq(leaveApply.ID)).Updates(leaveApply)
	if err != nil {
		response.Fail(ctx, "更新失败！")
		return
	}
	response.Success(ctx, "更新成功！")
}
