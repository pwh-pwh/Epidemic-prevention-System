package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/user_service"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"github.com/pwh-pwh/Epidemic-prevention-System/vo"
	"gorm.io/gen"
	"log"
	"strconv"
	"strings"
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

func ListUser(ctx *gin.Context) {
	var cds []gen.Condition
	userQ := query.Use(mysql.DB).SysUser
	userName := ctx.Query("userName")
	if userName != "" {
		cds = append(cds, userQ.Username.Like("%"+userName+"%"))
	}
	phoneNumber := ctx.Query("phoneNumber")
	if phoneNumber != "" {
		cds = append(cds, userQ.PhoneNumber.Like("%"+phoneNumber+"%"))
	}
	statusS := ctx.Query("status")
	if statusS != "" {
		status, _ := strconv.Atoi(statusS)
		cds = append(cds, userQ.Status.Eq(int32(status)))
	}
	deptIdS := ctx.Query("deptId")
	if deptIdS != "" {
		deptId, _ := strconv.Atoi(deptIdS)
		cds = append(cds, userQ.Status.Eq(int32(deptId)))
	}
	data, count, err := userQ.WithContext(context.Background()).Omit(userQ.Password).Where(cds...).FindByPage(utils.GetPage(ctx))
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, vo.PageVO{
		Records: data,
		Total:   count,
	})
}

func AddUser(ctx *gin.Context) {
	sysUser := new(models.SysUser)
	err := ctx.ShouldBindJSON(sysUser)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	sysUser.Password = utils.Md5Crypt(sysUser.Password)
	userQ := query.Use(mysql.DB).SysUser
	err = userQ.WithContext(context.Background()).Create(sysUser)
	if err != nil {
		response.Fail(ctx, "添加失败")
		return
	}
	response.Success(ctx, "添加成功")
}

func UpdateUser(ctx *gin.Context) {
	sysUser := new(models.SysUser)
	err := ctx.ShouldBindJSON(sysUser)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	err = mysql.DB.Model(sysUser).Select("*").Omit("create_time", "password").Where("id=?", sysUser.ID).Updates(sysUser).Error
	if err != nil {
		response.Fail(ctx, "修改失败")
		return
	}
	response.Success(ctx, "修改成功")
}

func DeleteUesr(ctx *gin.Context) {
	idsS := ctx.Query("ids")
	split := strings.Split(idsS, ",")
	var ids []int64
	for _, s := range split {
		parseInt, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, parseInt)
	}
	userQ := query.Use(mysql.DB).SysUser
	_, err := userQ.WithContext(context.Background()).Where(userQ.ID.In(ids...)).Delete()
	if err != nil {
		response.Fail(ctx, "删除失败")
		return
	}
	urQ := query.Use(mysql.DB).SysUserRole
	_, _ = urQ.WithContext(context.Background()).Where(urQ.UserID.In(ids...)).Delete()
	response.Success(ctx, "删除成功")
}

func ApplyUserRole(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, _ := strconv.Atoi(idS)
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	var roleIds []int64
	err := ctx.ShouldBindJSON(&roleIds)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	urQ := query.Use(mysql.DB).SysUserRole
	var sysUserRoles []*models.SysUserRole
	for _, roleId := range roleIds {
		userRole := new(models.SysUserRole)
		userRole.UserID = int64(id)
		userRole.RoleID = roleId
		sysUserRoles = append(sysUserRoles, userRole)
	}
	urQ.WithContext(context.Background()).Where(urQ.UserID.Eq(int64(id))).Delete()
	err = urQ.WithContext(context.Background()).Create(sysUserRoles...)
	if err != nil {
		response.Fail(ctx, "删除失败:"+err.Error())
		return
	}
	user_service.ClearUserAuthority(username)
	response.Success(ctx, "删除成功")
}

func ResetPwd(ctx *gin.Context) {
	sysUser := new(models.SysUser)
	err := ctx.ShouldBindJSON(sysUser)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	sysUser.Password = utils.Md5Crypt(sysUser.Password)
	err = mysql.DB.Model(sysUser).Select("*").Omit("create_time").Where("id=?", sysUser.ID).Updates(sysUser).Error
	if err != nil {
		response.Fail(ctx, "重置密码失败！")
		return
	}
	response.Success(ctx, "重置密码成功！")
}

func UpdateInfo(ctx *gin.Context) {
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	sysUser := new(models.SysUser)
	err := ctx.ShouldBindJSON(sysUser)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	sysUser.Password = utils.Md5Crypt(sysUser.Password)
	err = mysql.DB.Model(sysUser).Select("*").Select("city", "nickname", "phone_number", "remark").Where("id=?", sysUser.ID).Updates(sysUser).Error
	if err != nil {
		response.Fail(ctx, "更新个人信息失败")
		return
	}
	client := myredis.GetRedisClient()
	uQ := query.Use(mysql.DB).SysUser
	take, _ := uQ.WithContext(context.Background()).Where(uQ.ID.Eq(sysUser.ID)).Take()
	client.Set(myredis.UserPre+username, take, 36600*time.Second)
	response.Success(ctx, "更新个人信息成功")
}

func UpdatePassword(ctx *gin.Context) {
	oldPassword := ctx.Query("oldPassword")
	newPassword := ctx.Query("newPassword")
	confirmPassword := ctx.Query("confirmPassword")
	usernameI, _ := ctx.Get("username")
	username := usernameI.(string)
	if newPassword != confirmPassword {
		response.Fail(ctx, "两次输入的密码不一致！")
		return
	}
	userQ := query.Use(mysql.DB).SysUser
	sysUser, _ := userQ.WithContext(context.Background()).Where(userQ.Username.Eq(username)).Take()
	if utils.Md5Crypt(oldPassword) != sysUser.Password {
		response.Fail(ctx, "旧密码错误")
		return
	}
	sysUser.Password = utils.Md5Crypt(confirmPassword)
	_, err := userQ.WithContext(context.Background()).Select(userQ.Password).Where(userQ.ID.Eq(sysUser.ID)).Updates(sysUser)
	if err != nil {
		response.Fail(ctx, "修改密码失败")
		return
	}
	response.Success(ctx, "修改密码成功")
}

func Avatar(ctx *gin.Context) {
	file, err := ctx.FormFile("avatarFile")
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	userJson := ctx.PostForm("user")
	uploadFile, err := utils.UploadFile(file)
	if err != nil {
		response.Fail(ctx, "上传头像失败")
		return
	}
	sysUser := new(models.SysUser)
	err = json.Unmarshal([]byte(userJson), sysUser)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	userQ := query.Use(mysql.DB).SysUser
	sysUser.Avatar = uploadFile
	_, err = userQ.WithContext(context.Background()).Select(userQ.Avatar).Where(userQ.ID.Eq(sysUser.ID)).Updates(sysUser)
	if err != nil {
		response.Fail(ctx, "头像上传失败")
		return
	}
	client := myredis.GetRedisClient()
	client.Set(myredis.UserPre+sysUser.Username, sysUser, 36600*time.Second)
	response.Success(ctx, "头像上传成功")
}

func Info(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, _ := strconv.Atoi(idS)
	urQ := query.Use(mysql.DB).SysUserRole
	find, err := urQ.WithContext(context.Background()).Where(urQ.UserID.Eq(int64(id))).Find()
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, find)
}
