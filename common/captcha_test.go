package common

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func TestRedisCaptcha_Set(t *testing.T) {
	reidsClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "117.50.163.15", 6379),
	})
	captchaR := RedisCaptcha{reidsClient}
	captchaR.Set("11", []byte("aaa"))
}

func TestRedisCaptcha_Get(t *testing.T) {
	reidsClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "117.50.163.15", 6379),
	})
	captchaR := RedisCaptcha{reidsClient}
	digits := captchaR.Get("11", true)
	fmt.Println(string(digits))
}

func TestNewBase64Img(t *testing.T) {
	str, _ := NewBase64Img()
	fmt.Println(str)
}

func TestVerifyCaptcha(t *testing.T) {
	flag := VerifyCaptcha("gM3UJFSLrVsW106z2jsq", "982052")
	fmt.Println(flag)
}
