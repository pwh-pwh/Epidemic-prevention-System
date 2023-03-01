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
	err := db.AutoMigrate(&models.HealthClock{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&models.HealthClock{}) fail: %s", err)
	}
}

func TestSql(t *testing.T) {
	var count int
	db.Debug().Raw("select count(id) from health_clock where to_days(create_time) = to_days(now()) and username = ?", "admin").Scan(&count)
	fmt.Println(count)
}

func Test_healthClockQuery(t *testing.T) {
	healthClock := newHealthClock(db)
	healthClock = *healthClock.As(healthClock.TableName())
	_do := healthClock.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(healthClock.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <health_clock> fail:", err)
		return
	}

	_, ok := healthClock.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from healthClock success")
	}

	err = _do.Create(&models.HealthClock{})
	if err != nil {
		t.Error("create item in table <health_clock> fail:", err)
	}

	err = _do.Save(&models.HealthClock{})
	if err != nil {
		t.Error("create item in table <health_clock> fail:", err)
	}

	err = _do.CreateInBatches([]*models.HealthClock{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <health_clock> fail:", err)
	}

	_, err = _do.Select(healthClock.ALL).Take()
	if err != nil {
		t.Error("Take() on table <health_clock> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <health_clock> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <health_clock> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <health_clock> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*models.HealthClock{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <health_clock> fail:", err)
	}

	_, err = _do.Select(healthClock.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <health_clock> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <health_clock> fail:", err)
	}

	_, err = _do.Select(healthClock.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <health_clock> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <health_clock> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <health_clock> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <health_clock> fail:", err)
	}

	_, err = _do.ScanByPage(&models.HealthClock{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <health_clock> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <health_clock> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <health_clock> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <health_clock> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <health_clock> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <health_clock> fail:", err)
	}
}
