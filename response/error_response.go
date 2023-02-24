package response

import "errors"

var (
	NotValidErr        = errors.New("参数校验不通过")
	IllegalArgumentErr = errors.New("参数不合法")
	DuplicateKeyErr    = errors.New("重复添加")
)
