package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"log"
	"time"
)

func UserInfo(ctx *gin.Context) {
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	sysUser := new(models.SysUser)
	redisClient := myredis.GetRedisClient()
	err := redisClient.Get(myredis.UserPre + username).Scan(sysUser)
	if err != nil {
		log.Printf("redis get user error:%v\n", err)
		panic(err)
	}
	sysUser.Password = ""
	sysDept := new(models.SysDept)
	key := fmt.Sprintf(myredis.DeptPre+"%v", sysUser.DeptID)
	if result, _ := redisClient.Exists(key).Result(); result == 1 {
		err := redisClient.Get(key).Scan(sysDept)
		if err != nil {
			log.Printf("redis scan dept err :%v \n", err)
		}
	} else {
		deptQuery := query.Use(mysql.DB).SysDept
		sysDept, err = deptQuery.WithContext(context.Background()).Where(deptQuery.DeptID.Eq(sysUser.DeptID)).Take()
		if err != nil {
			log.Printf("dept dao get dept err :%v \n", err)
			panic(err)
		}
		redisClient.Set(key, sysDept, time.Hour)
	}
	response.Success(ctx, gin.H{
		"user": sysUser,
		"dept": sysDept,
	})
}
