package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/menu_service"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/user_service"
	"strconv"
)

func Nav(ctx *gin.Context) {
	auListI, ok := ctx.Get("auList")
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	var auList []string
	if ok {
		auList = auListI.([]string)
	} else {
		auList = user_service.GetUserAuthorityList(username)
	}
	//获取nav
	nav, err := menu_service.GetUserNav(username)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{
		"authoritys": auList,
		"nav":        nav,
	})
}

func ListMenu(ctx *gin.Context) {
	nav, err := menu_service.GetMenu()
	if err != nil {
		panic(err)
	}
	response.Success(ctx, nav)
}

func SaveMenu(ctx *gin.Context) {
	sysMenu := new(models.SysMenu)
	err := ctx.ShouldBindJSON(sysMenu)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	err = query.Use(mysql.DB).SysMenu.WithContext(context.Background()).Create(sysMenu)
	if err != nil {
		response.Fail(ctx, "添加失败")
		return
	}
	response.Success(ctx, "添加成功")
}

func EditMenu(ctx *gin.Context) {
	sysMenu := new(models.SysMenu)
	err := ctx.ShouldBindJSON(sysMenu)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	menuQ := query.Use(mysql.DB).SysMenu
	_, err = menuQ.WithContext(context.Background()).Select(menuQ.ParentID, menuQ.Status,
		menuQ.OrderNum, menuQ.Name, menuQ.Type, menuQ.Component, menuQ.Path, menuQ.Perms, menuQ.Icon).
		Where(menuQ.ID.Eq(sysMenu.ID)).Updates(sysMenu)
	if err != nil {
		response.Fail(ctx, "修改失败")
		return
	}
	user_service.ClearUserAuthorityByMenuId(sysMenu.ID)
	response.Success(ctx, "修改成功")
}

func DeleteMenu(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, _ := strconv.Atoi(idS)
	sysMenuQ := query.Use(mysql.DB).SysMenu
	count, err := sysMenuQ.WithContext(context.Background()).Where(sysMenuQ.ParentID.Eq(int64(id))).Count()
	if count > 0 {
		response.Fail(ctx, "请先删除下级数据")
		return
	}
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	_, err = sysMenuQ.WithContext(context.Background()).Where(sysMenuQ.ID.Eq(int64(id))).Delete()
	if err != nil {
		response.Fail(ctx, "删除失败")
		return
	}
	user_service.ClearUserAuthorityByMenuId(int64(id))
	roleMenuQ := query.Use(mysql.DB).SysRoleMenu
	roleMenuQ.WithContext(context.Background()).Where(roleMenuQ.MenuID.Eq(int64(id))).Delete()
	response.Success(ctx, "删除成功")
}
