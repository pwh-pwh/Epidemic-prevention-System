package service

import "github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"

//select count(id) from health_clock where to_days(create_time) = to_days(now()) and username = #{name}
func CheckClockToday(username string) int {
	var count int
	mysql.DB.Raw("select count(id) from health_clock where to_days(create_time) = to_days(now()) and username = ? and is_delete=0", username).Scan(&count)
	return count
}
