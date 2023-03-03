package menu_service

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/dto"
	"github.com/pwh-pwh/Epidemic-prevention-System/logic"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"log"
	"sort"
)

func ListByIds(ids []int64) ([]*models.SysMenu, error) {
	sysMenu := query.Use(mysql.DB).SysMenu
	return sysMenu.WithContext(context.Background()).Where(sysMenu.ID.In(ids...)).Find()
}
func GetMenu() ([]*models.SysMenu, error) {
	sysMenuQ := query.Use(mysql.DB).SysMenu
	menus, err := sysMenuQ.WithContext(context.Background()).Where().Order(sysMenuQ.OrderNum).Find()
	if err != nil {
		log.Printf("listbyids err:%v\n", err)
		return nil, err
	}
	treeMenu := buildTreeMenu(menus)
	return treeMenu, nil
}

func GetUserNav(username string) ([]*dto.NavMenu, error) {
	sysUser := new(models.SysUser)
	err := myredis.GetRedisClient().Get(myredis.UserPre + username).Scan(sysUser)
	if err != nil {
		log.Printf("redisclient get user err:%v\n", err)
		return nil, err
	}
	menuIds, err := logic.GetMenuIds(sysUser.ID)
	if err != nil {
		log.Printf("getmenuids err:%v\n", err)
		return nil, err
	}
	menus, err := ListByIds(menuIds)
	if err != nil {
		log.Printf("listbyids err:%v\n", err)
		return nil, err
	}
	treeMenu := buildTreeMenu(menus)
	return convert(treeMenu), nil
}

func buildTreeMenu(menus []*models.SysMenu) []*models.SysMenu {
	var finalMenus []*models.SysMenu
	for _, menu := range menus {
		for _, sysMenu := range menus {
			if menu.ID == sysMenu.ParentID {
				menu.Children = append(menu.Children, sysMenu)
			}
		}
		if menu.ParentID == 0 {
			finalMenus = append(finalMenus, menu)
		}
	}
	return finalMenus
}

func convert(menuTree []*models.SysMenu) []*dto.NavMenu {
	var navMenus []*dto.NavMenu
	for _, me := range menuTree {
		if me.Status != 0 {
			navMenu := dto.NavMenu{
				Id:        me.ID,
				Name:      me.Perms,
				Title:     me.Name,
				Icon:      me.Icon,
				Path:      me.Path,
				Component: me.Component,
				OrderNum:  int(me.OrderNum),
			}
			if len(me.Children) > 0 {
				navMenu.Children = convert(me.Children)
			}
			navMenus = append(navMenus, &navMenu)
		}
	}
	sort.Slice(navMenus, func(i, j int) bool {
		return navMenus[i].OrderNum < navMenus[j].OrderNum
	})
	for _, n := range navMenus {
		if len(n.Children) > 1 {
			sort.Slice(n.Children, func(i, j int) bool {
				return n.Children[i].OrderNum < n.Children[j].OrderNum
			})
		}
	}
	return navMenus
}
