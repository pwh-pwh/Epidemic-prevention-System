package service

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
)

func SaveLoginInfo(loginInfo *models.SysLoginInfo) {
	sysLoginInfo := query.Use(mysql.DB).SysLoginInfo
	sysLoginInfo.WithContext(context.Background()).Create(loginInfo)
}
