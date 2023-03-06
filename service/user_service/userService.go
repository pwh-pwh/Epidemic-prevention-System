package user_service

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/logic"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/menu_service"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"log"
	"strings"
	"time"
)

func GetByUserName(username string) *models.SysUser {
	user := query.Use(mysql.DB).SysUser
	take, err := user.WithContext(context.Background()).Where(user.Username.Eq(username), user.Status.Eq(1)).Take()
	if err != nil {
		log.Printf("GetByUserName error:%v", err)
		return nil
	}
	redisClient := myredis.GetRedisClient()
	redisClient.Set(myredis.UserPre+username, take, 60*60*time.Second)
	return take
}

func GetUserAuthorityList(username string) []string {
	redisClient := myredis.GetRedisClient()
	sysUser := new(models.SysUser)
	err := redisClient.Get(myredis.UserPre + username).Scan(sysUser)
	if err != nil {
		log.Printf("GetUseAuthority redis get sysUser error:%v\n", err)
	}
	var authorityList []string
	result, err := redisClient.Get(myredis.GrantedAuthorityPre + sysUser.Username).Result()
	if err == nil {
		authorityList = strings.Split(result, ",")
	} else {
		sysRoles, err := logic.GetRoleByUserId(sysUser.ID)
		if err != nil {
			log.Printf("GetuserAuL error :%v\n", err)
			return nil
		}
		for _, r := range sysRoles {
			authorityList = append(authorityList, "ROLE_"+r.Code)
		}
		menuIds, err := logic.GetMenuIds(sysUser.ID)
		if err != nil {
			log.Printf("get menuids error:%v\n", err)
		}
		if len(menuIds) > 0 {
			sysMenus, err := menu_service.ListByIds(menuIds)
			if err != nil {
				log.Printf("menu listbyids error:%v\n", err)
			}
			for _, me := range sysMenus {
				authorityList = append(authorityList, me.Perms)
			}
		}
		authorityStr := strings.Join(authorityList, ",")
		_, err = redisClient.Set(myredis.GrantedAuthorityPre+sysUser.Username, authorityStr, time.Hour).Result()
		if err != nil {
			log.Printf("redisclient set authoritylist err:%v\n", err)
		}
	}
	return authorityList
}

func RegisterUser(username, password, registerCode, phoneNumber string, roleType, deptId int) bool {
	roleId := utils.SwitchRole(roleType, registerCode)
	if roleId != -1 {
		user := new(models.SysUser)
		user.Username = username
		user.Password = utils.Md5Crypt(password)
		user.Avatar = common.DEFAULT_IMG
		user.DeptID = int64(deptId)
		user.PhoneNumber = phoneNumber
		user.Status = 1
		_ = query.Use(mysql.DB).SysUser.WithContext(context.Background()).Create(user)
		userRole := new(models.SysUserRole)
		userRole.UserID = user.ID
		userRole.RoleID = int64(roleId)
		_ = query.Use(mysql.DB).SysUserRole.WithContext(context.Background()).Create(userRole)
		return true
	}
	return false
}

//SELECT DISTINCT
//            su.*
//        FROM
//            sys_user_role ur
//        LEFT JOIN sys_role_menu rm ON ur.role_id = rm.role_id
//        LEFT JOIN sys_user su ON ur.user_id = su.id
//        WHERE
//            rm.menu_id = #{menuId}
//SELECT `sys_user`.`username` FROM `sys_user_role` LEFT JOIN `sys_role_menu` ON `sys_user_role`.`role_id` = `sys_role_menu`.`role_id`
//LEFT JOIN `sys_user` ON `sys_user`.`id` = `sys_user_role`.`user_id` WHERE `sys_role_menu`.`menu_id` = 25
func ClearUserAuthorityByMenuId(id int64) {
	sysUserQ := query.Use(mysql.DB).SysUser
	sysUserRoleQ := query.Use(mysql.DB).SysUserRole
	roleMenuQ := query.Use(mysql.DB).SysRoleMenu
	var usernameList []string
	sysUserRoleQ.WithContext(context.Background()).Distinct(sysUserQ.Username).Select(sysUserQ.Username).LeftJoin(roleMenuQ, sysUserRoleQ.RoleID.EqCol(roleMenuQ.RoleID)).
		LeftJoin(sysUserQ, sysUserQ.ID.EqCol(sysUserRoleQ.UserID)).Where(roleMenuQ.MenuID.Eq(id)).Scan(&usernameList)
	redisClient := myredis.GetRedisClient()
	for i, s := range usernameList {
		usernameList[i] = myredis.GrantedAuthorityPre + s
	}
	redisClient.Del(usernameList...)
}

/*
List<SysUser> sysUserList = list(new QueryWrapper<SysUser>()
                .inSql("id", "select user_id from sys_user_role where role_id = " + roleId));
        sysUserList.forEach(user -> clearUserAuthority(user.getUsername()));
*/
func ClearUserAuthorityByRoleId(roleId int64) {
	ctx := context.Background()
	sysUserQ := query.Use(mysql.DB).SysUser
	sysUserRoleQ := query.Use(mysql.DB).SysUserRole
	var usernameList []string
	urDo := sysUserRoleQ.WithContext(ctx).Select(sysUserRoleQ.UserID).Where(sysUserRoleQ.RoleID.Eq(roleId))
	sysUserQ.WithContext(ctx).Select(sysUserQ.Username).Where(
		sysUserQ.WithContext(ctx).Columns(sysUserQ.ID).In(urDo)).Scan(&usernameList)
	redisClient := myredis.GetRedisClient()
	for i, s := range usernameList {
		usernameList[i] = myredis.GrantedAuthorityPre + s
	}
	redisClient.Del(usernameList...)
}

func ClearUserAuthority(username string) {
	redisClient := myredis.GetRedisClient()
	redisClient.Del(myredis.GrantedAuthorityPre + username)
}
