package dto

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGEtServer(t *testing.T) {
	server := GetServer()
	marshal, err := json.Marshal(server)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("not error")
	fmt.Println(string(marshal))
}

func TestGetDirInfo(t *testing.T) {
	info := GetDirInfo()
	fmt.Println(len(info))
	for _, i := range info {
		fmt.Println(i)
	}
}
