// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:             db,
		AccessRegister: newAccessRegister(db, opts...),
		AccessReturn:   newAccessReturn(db, opts...),
		GoodInfo:       newGoodInfo(db, opts...),
		GoodStock:      newGoodStock(db, opts...),
		GoodType:       newGoodType(db, opts...),
		HealthClock:    newHealthClock(db, opts...),
		HealthReport:   newHealthReport(db, opts...),
		LeaveApply:     newLeaveApply(db, opts...),
		SysDept:        newSysDept(db, opts...),
		SysLoginInfo:   newSysLoginInfo(db, opts...),
		SysMenu:        newSysMenu(db, opts...),
		SysNotice:      newSysNotice(db, opts...),
		SysOperateLog:  newSysOperateLog(db, opts...),
		SysRole:        newSysRole(db, opts...),
		SysRoleMenu:    newSysRoleMenu(db, opts...),
		SysUser:        newSysUser(db, opts...),
		SysUserRole:    newSysUserRole(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AccessRegister accessRegister
	AccessReturn   accessReturn
	GoodInfo       goodInfo
	GoodStock      goodStock
	GoodType       goodType
	HealthClock    healthClock
	HealthReport   healthReport
	LeaveApply     leaveApply
	SysDept        sysDept
	SysLoginInfo   sysLoginInfo
	SysMenu        sysMenu
	SysNotice      sysNotice
	SysOperateLog  sysOperateLog
	SysRole        sysRole
	SysRoleMenu    sysRoleMenu
	SysUser        sysUser
	SysUserRole    sysUserRole
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		AccessRegister: q.AccessRegister.clone(db),
		AccessReturn:   q.AccessReturn.clone(db),
		GoodInfo:       q.GoodInfo.clone(db),
		GoodStock:      q.GoodStock.clone(db),
		GoodType:       q.GoodType.clone(db),
		HealthClock:    q.HealthClock.clone(db),
		HealthReport:   q.HealthReport.clone(db),
		LeaveApply:     q.LeaveApply.clone(db),
		SysDept:        q.SysDept.clone(db),
		SysLoginInfo:   q.SysLoginInfo.clone(db),
		SysMenu:        q.SysMenu.clone(db),
		SysNotice:      q.SysNotice.clone(db),
		SysOperateLog:  q.SysOperateLog.clone(db),
		SysRole:        q.SysRole.clone(db),
		SysRoleMenu:    q.SysRoleMenu.clone(db),
		SysUser:        q.SysUser.clone(db),
		SysUserRole:    q.SysUserRole.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		AccessRegister: q.AccessRegister.replaceDB(db),
		AccessReturn:   q.AccessReturn.replaceDB(db),
		GoodInfo:       q.GoodInfo.replaceDB(db),
		GoodStock:      q.GoodStock.replaceDB(db),
		GoodType:       q.GoodType.replaceDB(db),
		HealthClock:    q.HealthClock.replaceDB(db),
		HealthReport:   q.HealthReport.replaceDB(db),
		LeaveApply:     q.LeaveApply.replaceDB(db),
		SysDept:        q.SysDept.replaceDB(db),
		SysLoginInfo:   q.SysLoginInfo.replaceDB(db),
		SysMenu:        q.SysMenu.replaceDB(db),
		SysNotice:      q.SysNotice.replaceDB(db),
		SysOperateLog:  q.SysOperateLog.replaceDB(db),
		SysRole:        q.SysRole.replaceDB(db),
		SysRoleMenu:    q.SysRoleMenu.replaceDB(db),
		SysUser:        q.SysUser.replaceDB(db),
		SysUserRole:    q.SysUserRole.replaceDB(db),
	}
}

type queryCtx struct {
	AccessRegister *accessRegisterDo
	AccessReturn   *accessReturnDo
	GoodInfo       *goodInfoDo
	GoodStock      *goodStockDo
	GoodType       *goodTypeDo
	HealthClock    *healthClockDo
	HealthReport   *healthReportDo
	LeaveApply     *leaveApplyDo
	SysDept        *sysDeptDo
	SysLoginInfo   *sysLoginInfoDo
	SysMenu        *sysMenuDo
	SysNotice      *sysNoticeDo
	SysOperateLog  *sysOperateLogDo
	SysRole        *sysRoleDo
	SysRoleMenu    *sysRoleMenuDo
	SysUser        *sysUserDo
	SysUserRole    *sysUserRoleDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AccessRegister: q.AccessRegister.WithContext(ctx),
		AccessReturn:   q.AccessReturn.WithContext(ctx),
		GoodInfo:       q.GoodInfo.WithContext(ctx),
		GoodStock:      q.GoodStock.WithContext(ctx),
		GoodType:       q.GoodType.WithContext(ctx),
		HealthClock:    q.HealthClock.WithContext(ctx),
		HealthReport:   q.HealthReport.WithContext(ctx),
		LeaveApply:     q.LeaveApply.WithContext(ctx),
		SysDept:        q.SysDept.WithContext(ctx),
		SysLoginInfo:   q.SysLoginInfo.WithContext(ctx),
		SysMenu:        q.SysMenu.WithContext(ctx),
		SysNotice:      q.SysNotice.WithContext(ctx),
		SysOperateLog:  q.SysOperateLog.WithContext(ctx),
		SysRole:        q.SysRole.WithContext(ctx),
		SysRoleMenu:    q.SysRoleMenu.WithContext(ctx),
		SysUser:        q.SysUser.WithContext(ctx),
		SysUserRole:    q.SysUserRole.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
