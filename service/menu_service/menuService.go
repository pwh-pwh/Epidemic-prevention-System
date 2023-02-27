package menu_service

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
)

func ListByIds(ids []int64) ([]*models.SysMenu, error) {
	sysMenu := query.Use(mysql.DB).SysMenu
	return sysMenu.WithContext(context.Background()).Where(sysMenu.ID.In(ids...)).Find()
}
