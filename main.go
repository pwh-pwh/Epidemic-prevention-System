package main

import (
	"fmt"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/routers"
	"github.com/pwh-pwh/Epidemic-prevention-System/settings"
	"time"
)

// @title 校园疫情管理系统
// @version 1.0
// @description 校园疫情管理系统
// @contact.name coderpwh
// @contact.email plmoknpwh@163.com
func main() {
	common.StartTime = time.Now()
	common.InitAll()
	// 注册路由
	r := routers.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
