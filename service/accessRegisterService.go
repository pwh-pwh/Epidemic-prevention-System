package service

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
)

func AddAccessRegister(ar *models.AccessRegister) bool {
	register := query.Use(mysql.DB).AccessRegister
	accessReturn := query.Use(mysql.DB).AccessReturn
	err := register.WithContext(context.Background()).Create(ar)
	if err != nil {
		return false
	}
	if ar.Type == 1 {
		areturn := new(models.AccessReturn)
		areturn.Name = ar.Name
		areturn.Card = ar.Card
		areturn.Dept = ar.Dept
		areturn.Phone = ar.Phone
		areturn.Remark = ar.Remark
		accessReturn.WithContext(context.Background()).Create(areturn)
	} else {
		accessReturn.WithContext(context.Background()).Where(accessReturn.Name.Eq(ar.Name)).Delete()
	}
	return true
}
