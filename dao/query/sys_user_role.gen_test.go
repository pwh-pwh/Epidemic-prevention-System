// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"testing"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&models.SysUserRole{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&models.SysUserRole{}) fail: %s", err)
	}
}

func Test_sysUserRoleQuery(t *testing.T) {
	sysUserRole := newSysUserRole(db)
	sysUserRole = *sysUserRole.As(sysUserRole.TableName())
	_do := sysUserRole.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(sysUserRole.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <sys_user_role> fail:", err)
		return
	}

	_, ok := sysUserRole.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from sysUserRole success")
	}

	err = _do.Create(&models.SysUserRole{})
	if err != nil {
		t.Error("create item in table <sys_user_role> fail:", err)
	}

	err = _do.Save(&models.SysUserRole{})
	if err != nil {
		t.Error("create item in table <sys_user_role> fail:", err)
	}

	err = _do.CreateInBatches([]*models.SysUserRole{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <sys_user_role> fail:", err)
	}

	_, err = _do.Select(sysUserRole.ALL).Take()
	if err != nil {
		t.Error("Take() on table <sys_user_role> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <sys_user_role> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*models.SysUserRole{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Select(sysUserRole.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Select(sysUserRole.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <sys_user_role> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <sys_user_role> fail:", err)
	}

	_, err = _do.ScanByPage(&models.SysUserRole{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <sys_user_role> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <sys_user_role> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <sys_user_role> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <sys_user_role> fail:", err)
	}
}