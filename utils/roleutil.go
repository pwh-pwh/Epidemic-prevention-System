package utils

import "github.com/pwh-pwh/Epidemic-prevention-System/common"

func SwitchRole(roleType int, registerCode string) int {
	var r int
	switch roleType {
	case common.ROLE_STUDENT:
		if registerCode == common.CodeConfig.StudentCode {
			r = common.CodeConfig.StudentRole
		} else {
			r = -1
		}
	case common.ROLE_TEACHER:
		if registerCode == common.CodeConfig.TeacherCode {
			r = common.CodeConfig.TeacherRole
		} else {
			r = -1
		}
	case common.ROLE_SERVICE:
		if registerCode == common.CodeConfig.ServiceCode {
			r = common.CodeConfig.ServiceRole
		} else {
			r = -1
		}
	default:
		r = -1
	}
	return r
}
