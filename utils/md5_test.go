package utils

import (
	"fmt"
	"testing"
)

func TestMd5Crypt(t *testing.T) {
	str := Md5Crypt("123456")
	fmt.Println(str)
}
