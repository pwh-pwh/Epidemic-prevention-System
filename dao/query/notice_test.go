package query

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"log"
	"testing"
)

func TestSysNotice(t *testing.T) {
	notice := Use(db).SysNotice
	m := models.SysNotice{
		Title:    "t1",
		Content:  "t2c",
		Status:   0,
		IsDelete: 0,
		Remark:   "",
	}
	err := notice.WithContext(context.Background()).Debug().Create(&m)
	if err != nil {
		log.Println(err)
	}
}
