// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"encoding"
	"encoding/json"
	"gorm.io/plugin/soft_delete"
)

const TableNameSysNotice = "sys_notice"

// SysNotice mapped from table <sys_notice>
type SysNotice struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title      string                `gorm:"column:title" json:"title"`
	Content    string                `gorm:"column:content" json:"content"`
	Status     int32                 `gorm:"column:status;not null" json:"status"`
	CreateBy   string                `gorm:"column:create_by;type:varchar(50)" json:"create_by"`                     // 创建人
	UpdateBy   string                `gorm:"column:create_by;type:varchar(50)" json:"update_by"`                     // 创建人
	CreateTime LocalTime             `gorm:"column:create_time;type:datetime;autoCreateTime:true" json:"createTime"` // 创建时间
	UpdateTime LocalTime             `gorm:"column:update_time;type:datetime;autoUpdateTime:true" json:"updateTime"` // 更新时间
	IsDelete   soft_delete.DeletedAt `gorm:"column:is_delete;type:int(1);softDelete:flag" json:"-"`                  // 逻辑删除
	Remark     string                `gorm:"column:remark" json:"remark"`
}

func (n *SysNotice) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *SysNotice) MarshalBinary() (data []byte, err error) {
	return json.Marshal(n)
}

// TableName SysNotice's table name
func (*SysNotice) TableName() string {
	return TableNameSysNotice
}

var _ encoding.BinaryMarshaler = new(SysNotice)
var _ encoding.BinaryUnmarshaler = new(SysNotice)