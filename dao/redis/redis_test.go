package redis

import (
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

var _ encoding.BinaryMarshaler = &Aa{}
var _ encoding.BinaryUnmarshaler = &Aa{}

type Aa struct {
	Name string
	Age  int
}

func (a *Aa) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

func (a *Aa) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func TestGetRedisClient(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "117.50.163.15", 6379),
	})
	a := Aa{
		"coderpwh",
		11,
	}

	result, err := redisClient.Set("aat", &a, time.Hour).Result()
	fmt.Printf("result:%v err:%v\n", result, err)
	as := &Aa{}
	redisClient.Get("aat").Scan(as)
	fmt.Println(as)
}

func TestKeyNil(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "117.50.163.15", 6379),
	})
	result, err := redisClient.Get("faef").Result()
	fmt.Printf("result:%v err:%v\n", result, err)
}
