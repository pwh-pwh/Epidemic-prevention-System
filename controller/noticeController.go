package controller

import (
	"github.com/gin-gonic/gin"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
)

/**
@GetMapping
    public Result getNotice(){
        if (redisUtil.hasKey(Const.NOTICE_KEY)){
            SysNotice sysNotice = (SysNotice) redisUtil.get(Const.NOTICE_KEY);
            return Result.succ(sysNotice);
        }
        return Result.succ("暂无公告");
    }
*/

func GetNotice(ctx *gin.Context) {
	redisClient := myredis.GetRedisClient()
	if exist := redisClient.Exists(myredis.NoticeKey).Val(); exist == 1 {
		sysNotice := new(models.SysNotice)
		_ = redisClient.Get(myredis.NoticeKey).Scan(sysNotice)
		response.Success(ctx, sysNotice)
		return
	}
	response.Success(ctx, "暂无公告")
}
