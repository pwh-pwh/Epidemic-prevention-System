package query

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestUseDao(t *testing.T) {
	sysRole := Use(db).SysRole
	sysUserRole := Use(db).SysUserRole
	//SELECT id,name,code,remark,create_time,update_time,status,is_delete,version
	//FROM sys_role WHERE is_delete=0 AND (id IN (select role_id from sys_user_role where user_id = 1) AND status = ?)
	userRoleDo := sysUserRole.WithContext(context.Background()).Select(sysUserRole.RoleID).Where(sysUserRole.UserID.Eq(1))
	sysRoles, err := sysRole.WithContext(context.Background()).Where(sysUserRole.WithContext(context.Background()).Columns(sysRole.ID).In(userRoleDo)).Find()
	if err != nil {
		log.Printf("GetuserAuL error :%v", err)
	}
	for _, r := range sysRoles {
		fmt.Println(r)
	}
	sysRoleMenu := Use(db).SysRoleMenu
	var menuIds []int
	err = sysUserRole.WithContext(context.Background()).Debug().Distinct(sysRoleMenu.MenuID).LeftJoin(sysRoleMenu, sysUserRole.RoleID.EqCol(sysRoleMenu.RoleID)).Where(sysUserRole.UserID.Eq(1)).Scan(&menuIds)
	if err != nil {
		log.Printf("get menuids error:%v", err)
	}
	for _, me := range menuIds {
		fmt.Println(me)
	}
}
