package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	s := new(SysUser)
	marshal, _ := json.Marshal(s)
	fmt.Println(string(marshal))
	err := json.Unmarshal(marshal, &s)
	fmt.Println(err)
	fmt.Println(s)
}
