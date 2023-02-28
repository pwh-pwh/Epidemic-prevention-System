package common

import (
	"fmt"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/settings"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func loadConfig() {
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
}

func InitAll() {
	loadConfig()
	if err := mysql.InitializeDB(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	//初始化redis
	if err := redis.InitRedis(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	if err := InitializeJWT(settings.Conf.JwtConfig); err != nil {
		fmt.Printf("init jwt config failed, err:%v\n", err)
		return
	}
	initRedisCatpcha()
}
