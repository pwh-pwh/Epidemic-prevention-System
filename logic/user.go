package logic

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
)

func GetRoleByUserId(userId int64) ([]*models.SysRole, error) {
	sysRole := query.Use(mysql.DB).SysRole
	sysUserRole := query.Use(mysql.DB).SysUserRole
	//SELECT * FROM `sys_role` WHERE `sys_role`.`id` IN (SELECT `sys_user_role`.`role_id` FROM `sys_user_role` WHERE `sys_user_role`.`user_id` = 1) AND `sys_role`.`is_delete` = 0
	userRoleDo := sysUserRole.WithContext(context.Background()).Select(sysUserRole.RoleID).Where(sysUserRole.UserID.Eq(userId))
	return sysRole.WithContext(context.Background()).Where(sysUserRole.WithContext(context.Background()).Columns(sysRole.ID).In(userRoleDo)).Find()
}

func GetMenuIds(userId int64) ([]int64, error) {
	sysUserRole := query.Use(mysql.DB).SysUserRole
	sysRoleMenu := query.Use(mysql.DB).SysRoleMenu
	var menuIds []int64
	//getmenuids
	// SELECT DISTINCT `sys_role_menu`.`menu_id` FROM `sys_user_role` LEFT JOIN `sys_role_menu` ON `sys_user_role`.`role_id` = `sys_role_menu`.`role_id` WHERE `sys_user_role`.`user_id` = 1
	err := sysUserRole.WithContext(context.Background()).Distinct(sysRoleMenu.MenuID).LeftJoin(sysRoleMenu, sysUserRole.RoleID.EqCol(sysRoleMenu.RoleID)).Where(sysUserRole.UserID.Eq(userId)).Scan(&menuIds)
	return menuIds, err
}
