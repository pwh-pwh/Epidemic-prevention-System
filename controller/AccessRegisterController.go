package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
)

func GetAccessRegisterList(c *gin.Context) {
	accessRegister := query.Use(mysql.DB).AccessRegister
	find, err := accessRegister.WithContext(context.Background()).Find()
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": find,
	})
}
