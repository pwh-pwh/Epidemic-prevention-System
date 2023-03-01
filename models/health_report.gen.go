// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameHealthReport = "health_report"

// HealthReport mapped from table <health_report>
type HealthReport struct {
	ID          int64     `gorm:"column:id;type:bigint(11);primaryKey;autoIncrement:true" json:"id"`      // 自增id
	Username    string    `gorm:"column:username;type:varchar(50)" json:"username"`                       // 用户名
	DeptID      int64     `gorm:"column:dept_id;type:bigint(11)" json:"deptId"`                           // 部门id
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(50)" json:"phoneNumber"`                // 手机号
	Img3        string    `gorm:"column:img3;type:varchar(255)" json:"img3"`                              // 核酸报告
	Img2        string    `gorm:"column:img2;type:varchar(255)" json:"img2"`                              // 行程码
	Img1        string    `gorm:"column:img1;type:varchar(255)" json:"img1"`                              // 健康码
	Type        int32     `gorm:"column:type;type:int(1)" json:"type"`                                    // 返校情况
	CreateTime  LocalTime `gorm:"column:create_time;type:datetime;autoCreateTime:true" json:"createTime"` // 创建时间
	UpdateTime  LocalTime `gorm:"column:update_time;type:datetime;autoUpdateTime:true" json:"updateTime"` // 更新时间
}

// TableName HealthReport's table name
func (*HealthReport) TableName() string {
	return TableNameHealthReport
}
