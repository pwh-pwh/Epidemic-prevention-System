package user_service

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"log"
)

func GetByUserName(username string) *models.SysUser {
	user := query.Use(mysql.DB).SysUser
	take, err := user.WithContext(context.Background()).Where(user.Username.Eq(username), user.Status.Eq(1)).Take()
	if err != nil {
		log.Printf("GetByUserName error:%v", err)
		return nil
	}
	return take
}
