package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"strconv"
	"strings"
	"time"
)

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

func ListNotice(ctx *gin.Context) {
	var cds []gen.Condition
	sysNoticeQ := query.Use(mysql.DB).SysNotice
	title := ctx.Query("title")
	if title != "" {
		cds = append(cds, sysNoticeQ.Title.Like("%"+title+"%"))
	}
	createBy := ctx.Query("createBy")
	if createBy != "" {
		cds = append(cds, sysNoticeQ.CreateBy.Like("%"+createBy+"%"))
	}
	start := ctx.Query("start")
	end := ctx.Query("end")
	if start != "" && end != "" {
		sT, err := utils.ParseTime(start+" 00:00:00", common.TimeFormat)
		if err != nil {
			response.Fail(ctx, "时间参数错误")
			return
		}
		eT, err := utils.ParseTime(end+" 23:59:59", common.TimeFormat)
		if err != nil {
			response.Fail(ctx, "时间参数错误")
			return
		}
		cds = append(cds, sysNoticeQ.CreateTime.Between(sT, eT))
	}
	offset, limit := utils.GetPage(ctx)
	data, count, err := sysNoticeQ.WithContext(context.Background()).Where(cds...).Order(sysNoticeQ.CreateTime.Desc()).FindByPage(offset, limit)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func SaveNotice(ctx *gin.Context) {
	sysNotice := new(models.SysNotice)
	err := ctx.ShouldBindJSON(sysNotice)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	sysNotice.CreateBy = username
	err = query.Use(mysql.DB).SysNotice.WithContext(context.Background()).Create(sysNotice)
	if err != nil {
		response.Fail(ctx, "新增公告失败")
		return
	}
	response.Success(ctx, "新增公告成功")
}

func UpdateNotice(ctx *gin.Context) {
	sysNotice := new(models.SysNotice)
	err := ctx.ShouldBindJSON(sysNotice)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	sysNotice.UpdateBy = username
	noticeQ := query.Use(mysql.DB).SysNotice
	_, err = noticeQ.WithContext(context.Background()).Select(noticeQ.UpdateBy, noticeQ.Title, noticeQ.Status, noticeQ.Remark, noticeQ.Content).Where(noticeQ.ID.Eq(sysNotice.ID)).Updates(sysNotice)
	if err != nil {
		response.Fail(ctx, "更新公告失败")
		return
	}
	response.Success(ctx, "更新公告成功")
}
func DeleteNotice(ctx *gin.Context) {
	idsS := ctx.Query("ids")
	split := strings.Split(idsS, ",")
	var ids []int64
	for _, s := range split {
		parseInt, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, parseInt)
	}
	noticeQ := query.Use(mysql.DB).SysNotice
	_, err := noticeQ.WithContext(context.Background()).Where(noticeQ.ID.In(ids...)).Delete()
	if err != nil {
		response.Fail(ctx, "删除失败")
		return
	}
	response.Success(ctx, "删除成功")
}

func SetNotice(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, _ := strconv.Atoi(idS)
	noticeQ := query.Use(mysql.DB).SysNotice
	notice, _ := noticeQ.WithContext(context.Background()).Where(noticeQ.ID.Eq(int64(id))).Take()
	uid := "aa00bb"
	notice.NoticeId = uid
	redisClient := myredis.GetRedisClient()
	redisClient.Set(myredis.NoticeKey, notice, time.Second*86400)
	response.Success(ctx, "公告设置成功")
}
