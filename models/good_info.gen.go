// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameGoodInfo = "good_info"

// GoodInfo mapped from table <good_info>
type GoodInfo struct {
	ID         int64     `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"` // 物资信息id
	TypeID     int64     `gorm:"column:type_id;type:bigint(20)" json:"type_id"`                     // 类型id
	Name       string    `gorm:"column:name;type:varchar(50)" json:"name"`                          // 物资名称
	Img        string    `gorm:"column:img;type:varchar(255)" json:"img"`                           // 图片链接
	Size       string    `gorm:"column:size;type:varchar(50)" json:"size"`                          // 物资规格
	Unit       string    `gorm:"column:unit;type:varchar(50)" json:"unit"`                          // 物资单位
	CreateBy   string    `gorm:"column:create_by;type:varchar(50)" json:"create_by"`                // 创建人
	Remark     string    `gorm:"column:remark;type:varchar(50)" json:"remark"`                      // 备注
	Total      int32     `gorm:"column:total;type:int(11)" json:"total"`                            // 库存
	Status     int32     `gorm:"column:status;type:int(1);default:1" json:"status"`                 // 是否启用
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`               // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`               // 更新时间
	IsDelete   int32     `gorm:"column:is_delete;type:int(1)" json:"is_delete"`                     // 逻辑删除
	Version    int32     `gorm:"column:version;type:int(11);default:1" json:"version"`              // 乐观锁
}

// TableName GoodInfo's table name
func (*GoodInfo) TableName() string {
	return TableNameGoodInfo
}