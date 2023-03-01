package service

import "github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"

//select count(id) from health_report where to_days(create_time) = to_days(now()) and username = #{name}
func CheckReportToday(username string) int {
	var count int
	mysql.DB.Raw("select count(id) from health_report where to_days(create_time) = to_days(now()) and username = ? ", username).Scan(&count)
	return count
}
