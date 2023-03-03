package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/user_service"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
	"strings"
)

func InfoRole(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, _ := strconv.Atoi(idS)
	roleQ := query.Use(mysql.DB).SysRole
	sysRole, _ := roleQ.WithContext(context.Background()).Where(roleQ.ID.Eq(int64(id))).Take()
	menuQ := query.Use(mysql.DB).SysRoleMenu
	menuQ.WithContext(context.Background()).Select(menuQ.MenuID).Where(menuQ.RoleID.Eq(int64(id))).Scan(&(sysRole.MenuIds))
	response.Success(ctx, sysRole)
}

func ListRole(ctx *gin.Context) {
	var cds []gen.Condition
	roleQ := query.Use(mysql.DB).SysRole
	roleName := ctx.Query("roleName")
	if roleName != "" {
		cds = append(cds, roleQ.Name.Like("%"+roleName+"%"))
	}
	roleKey := ctx.Query("roleKey")
	if roleKey != "" {
		cds = append(cds, roleQ.Code.Like("%"+roleKey+"%"))
	}
	statusS := ctx.Query("status")
	if statusS != "" {
		status, _ := strconv.Atoi(statusS)
		cds = append(cds, roleQ.Status.Eq(int32(status)))
	}
	data, count, err := roleQ.WithContext(context.Background()).Where(cds...).FindByPage(utils.GetPage(ctx))
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func AddRole(ctx *gin.Context) {
	sysRole := new(models.SysRole)
	err := ctx.ShouldBindJSON(sysRole)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	roleQ := query.Use(mysql.DB).SysRole
	err = roleQ.WithContext(context.Background()).Create(sysRole)
	if err != nil {
		response.Fail(ctx, "添加失败")
		return
	}
	response.Success(ctx, "添加成功")
}

func EditRole(ctx *gin.Context) {
	sysRole := new(models.SysRole)
	err := ctx.ShouldBindJSON(sysRole)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	roleQ := query.Use(mysql.DB).SysRole
	_, err = roleQ.WithContext(context.Background()).Select(roleQ.Status, roleQ.Name, roleQ.Code, roleQ.Remark).Where(roleQ.ID.Eq(sysRole.ID)).Updates(sysRole)
	if err != nil {
		response.Fail(ctx, "修改失败")
		return
	}
	roleMenuQ := query.Use(mysql.DB).SysRoleMenu
	_, _ = roleMenuQ.WithContext(context.Background()).Where(roleMenuQ.RoleID.Eq(sysRole.ID)).Delete()
	menuIds := sysRole.MenuIds
	var roleMenus []*models.SysRoleMenu
	for _, menuId := range menuIds {
		roleMenu := new(models.SysRoleMenu)
		roleMenu.RoleID = sysRole.ID
		roleMenu.MenuID = menuId
		roleMenus = append(roleMenus, roleMenu)
	}
	roleMenuQ.WithContext(context.Background()).Create(roleMenus...)
	user_service.ClearUserAuthorityByRoleId(sysRole.ID)
	response.Success(ctx, "修改成功")
}

func DeleteRole(ctx *gin.Context) {
	idsS := ctx.Query("ids")
	split := strings.Split(idsS, ",")
	var ids []int64
	for _, s := range split {
		parseInt, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, parseInt)
	}
	roleQ := query.Use(mysql.DB).SysRole
	_, err := roleQ.WithContext(context.Background()).Where(roleQ.ID.In(ids...)).Delete()
	if err != nil {
		response.Fail(ctx, "删除失败")
		return
	}
	for _, id := range ids {
		user_service.ClearUserAuthorityByRoleId(id)
	}
	urQ := query.Use(mysql.DB).SysUserRole
	urQ.WithContext(context.Background()).Where(urQ.RoleID.In(ids...)).Delete()
	rmQ := query.Use(mysql.DB).SysRoleMenu
	rmQ.WithContext(context.Background()).Where(rmQ.RoleID.In(ids...)).Delete()
	response.Success(ctx, "删除成功")
}
