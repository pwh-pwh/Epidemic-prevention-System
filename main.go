package main

import (
	"fmt"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.InitializeDB(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	var ars []models.AccessRegister
	mysql.DB.Find(&ars)
	for _, item := range ars {
		fmt.Println(item)
	}
}
