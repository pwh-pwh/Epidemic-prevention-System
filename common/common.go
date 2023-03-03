package common

import (
	"fmt"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/settings"
	"time"
)

const (
	TimeFormat       = "2006-01-02 15:04:05"
	RoleTeacher      = "ROLE_teacher"
	ROLE_STUDENT     = 2
	ROLE_TEACHER     = 1
	ROLE_SERVICE     = 0
	DEFAULT_IMG      = "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
	DEFAULT_PASSWORD = "123456"
)

var StartTime time.Time

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
	InitOss(settings.Conf.OssConfig)
	InitCode(settings.Conf.CodeConfig)
}

func InitOss(cfg *settings.OssConfig) {
	Oss = cfg
}

func InitCode(cfg *settings.CodeConfig) {
	CodeConfig = cfg
}
