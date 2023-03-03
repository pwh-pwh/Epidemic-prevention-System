package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service"
	"strconv"
)

func ListDept(ctx *gin.Context) {
	flagS := ctx.Param("flag")
	depts, err := service.GetAllDept(flagS)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, depts)
}

func SaveDept(ctx *gin.Context) {
	sysDept := new(models.SysDept)
	err := ctx.ShouldBindJSON(sysDept)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	sysDept.CreateBy = username
	err = query.Use(mysql.DB).SysDept.WithContext(context.Background()).Create(sysDept)
	if err != nil {
		response.Fail(ctx, "添加失败！")
		return
	}
	response.Success(ctx, "添加成功！")
}

func UpdateDept(ctx *gin.Context) {
	sysDept := new(models.SysDept)
	err := ctx.ShouldBindJSON(sysDept)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	sysDept.UpdateBy = username
	sysDeptQ := query.Use(mysql.DB).SysDept
	_, err = sysDeptQ.WithContext(context.Background()).Select(sysDeptQ.DeptName, sysDeptQ.Status,
		sysDeptQ.ParentID, sysDeptQ.OrderNum, sysDeptQ.Phone, sysDeptQ.Email, sysDeptQ.Leader, sysDeptQ.UpdateBy).
		Where(sysDeptQ.DeptID.Eq(sysDept.DeptID)).Updates(sysDept)
	if err != nil {
		response.Fail(ctx, "修改失败!")
		return
	}
	response.Success(ctx, "修改成功!")
}

func DeleteDeptById(ctx *gin.Context) {
	id := ctx.Param("id")
	idNum, _ := strconv.Atoi(id)
	sysDeptQ := query.Use(mysql.DB).SysDept
	_, err := sysDeptQ.WithContext(context.Background()).Where(sysDeptQ.DeptID.Eq(int64(idNum))).Delete()
	if err != nil {
		response.Fail(ctx, "删除失败!")
		return
	}
	sysUserQ := query.Use(mysql.DB).SysUser
	_, err = sysUserQ.WithContext(context.Background()).Where(sysUserQ.DeptID.Eq(int64(idNum))).UpdateSimple(sysUserQ.DeptID.Value(100))
	if err != nil {
		response.Fail(ctx, err.Error())
	}
	response.Success(ctx, "删除成功!")
}
