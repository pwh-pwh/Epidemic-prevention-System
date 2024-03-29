// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameSysOperateLog = "sys_operate_log"

// SysOperateLog mapped from table <sys_operate_log>
type SysOperateLog struct {
	ID            int64     `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"` // 日志主键
	Title         string    `gorm:"column:title;type:varchar(50)" json:"title"`                        // 模块标题
	BusinessType  string    `gorm:"column:business_type;type:varchar(50)" json:"businessType"`         // 业务类型
	Method        string    `gorm:"column:method;type:varchar(100)" json:"method"`                     // 方法名称
	RequestMethod string    `gorm:"column:request_method;type:varchar(10)" json:"requestMethod"`       // 请求方式
	OperURL       string    `gorm:"column:oper_url;type:varchar(255)" json:"operUrl"`                  // 请求URL
	OperIP        string    `gorm:"column:oper_ip;type:varchar(128)" json:"operIp"`                    // 主机地址
	OperLocation  string    `gorm:"column:oper_location;type:varchar(255)" json:"operLocation"`        // 操作地点
	OperParam     string    `gorm:"column:oper_param;type:varchar(2000)" json:"operParam"`             // 请求参数
	OperName      string    `gorm:"column:oper_name;type:varchar(50)" json:"operName"`                 // 操作人
	JSONResult    string    `gorm:"column:json_result;type:varchar(2000)" json:"jsonResult"`           // 返回参数
	Status        int32     `gorm:"column:status;type:int(1);default:1" json:"status"`                 // 操作状态（1正常 0异常）
	ErrorMsg      string    `gorm:"column:error_msg;type:varchar(2000)" json:"errorMsg"`               // 错误消息
	OperTime      LocalTime `gorm:"column:oper_time;type:datetime" json:"operTime"`                    // 操作时间
}

// TableName SysOperateLog's table name
func (*SysOperateLog) TableName() string {
	return TableNameSysOperateLog
}
