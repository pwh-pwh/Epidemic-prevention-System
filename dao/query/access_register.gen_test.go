// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"log"
	"testing"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&models.AccessRegister{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&models.AccessRegister{}) fail: %s", err)
	}
}

func Test_GetARList(t *testing.T) {
	register := Use(db).AccessRegister
	find, err := register.WithContext(context.Background()).Find()
	if err != nil {
		log.Println(err)
	}
	for _, i := range find {
		log.Println(i)
	}
}

func Test_ArUpdate(t *testing.T) {
	register := Use(db).AccessRegister
	accessRegister := models.AccessRegister{
		ID:     1,
		Remark: "正常返校1",
	}
	_, err := register.WithContext(context.Background()).Where(register.ID.Eq(accessRegister.ID)).Update(register.Remark, accessRegister.Remark)
	if err != nil {
		log.Println(err)
	}
}

func Test_ArDelete(t *testing.T) {
	register := Use(db).AccessRegister
	_, err := register.WithContext(context.Background()).Debug().Where(register.ID.Eq(1)).Delete()
	if err != nil {
		log.Println(err)
	}
}

func Test_ArCreate(t *testing.T) {
	register := Use(db).AccessRegister
	accessRegister := models.AccessRegister{
		Name:   "李文",
		Phone:  "13226948870",
		Type:   1,
		Card:   "342623199906214412",
		Remark: "正常返校",
		Dept:   "软件186",
	}
	err := register.WithContext(context.Background()).Debug().Create(&accessRegister)
	if err != nil {
		log.Println(err)
	}
}

func Test_accessRegisterQuery(t *testing.T) {
	accessRegister := newAccessRegister(db)
	accessRegister = *accessRegister.As(accessRegister.TableName())
	_do := accessRegister.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(accessRegister.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <access_register> fail:", err)
		return
	}

	_, ok := accessRegister.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from accessRegister success")
	}

	err = _do.Create(&models.AccessRegister{})
	if err != nil {
		t.Error("create item in table <access_register> fail:", err)
	}

	err = _do.Save(&models.AccessRegister{})
	if err != nil {
		t.Error("create item in table <access_register> fail:", err)
	}

	err = _do.CreateInBatches([]*models.AccessRegister{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <access_register> fail:", err)
	}

	_, err = _do.Select(accessRegister.ALL).Take()
	if err != nil {
		t.Error("Take() on table <access_register> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <access_register> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <access_register> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <access_register> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*models.AccessRegister{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <access_register> fail:", err)
	}

	_, err = _do.Select(accessRegister.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <access_register> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <access_register> fail:", err)
	}

	_, err = _do.Select(accessRegister.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <access_register> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <access_register> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <access_register> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <access_register> fail:", err)
	}

	_, err = _do.ScanByPage(&models.AccessRegister{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <access_register> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <access_register> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <access_register> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <access_register> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <access_register> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <access_register> fail:", err)
	}
}