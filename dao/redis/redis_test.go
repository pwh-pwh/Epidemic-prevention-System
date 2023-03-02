package redis

import (
	"bufio"
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"strings"
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

func TestHasKey(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "117.50.163.15", 6379),
	})
	result, err := redisClient.Exists("aa").Result()
	fmt.Printf("result :%v err:%v", result, err)
}

func TestRedisInfo(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "117.50.163.15", 6379),
	})
	infoMap := make(map[string]string, 20)
	s := redisClient.Info().String()
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		text := sc.Text()
		if strings.Contains(text, ":") {
			split := strings.Split(text, ":")
			infoMap[split[0]] = split[1]
		}
	}
	marshal, _ := json.Marshal(infoMap)
	fmt.Println(string(marshal))
}

func TestCommandstats(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "117.50.163.15", 6379),
	})
	cs := redisClient.Info("commandstats").String()
	csBf := bufio.NewScanner(strings.NewReader(cs))
	csList := make([]nameValue, 0, 20)
	isFirst := true
	for csBf.Scan() {
		text := csBf.Text()
		if isFirst {
			isFirst = false
			continue
		}
		if !strings.Contains(text, ":") {
			continue
		}
		split := strings.Split(text, ":")
		csList = append(csList, nameValue{
			Name:  RemoveStart(split[0], "cmdstat_"),
			Value: SubstringBetween(split[1], "calls=", ",usec="),
		})
	}
	marshal, _ := json.Marshal(csList)
	fmt.Println(string(marshal))
}

type nameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func SubstringBetween(str string, open string, close string) string {
	start := strings.Index(str, open)
	if start != -1 {
		end := strings.Index(str, close)
		if end != -1 {
			return str[start+len(open) : end]
		}
	}
	return ""
}

func RemoveStart(str string, remove string) string {
	start := strings.Index(str, remove)
	if start != -1 {
		return str[start+len(remove):]
	}
	return str
}

func TestSubB(t *testing.T) {
	fmt.Println(SubstringBetween("cmdstat_exists:calls=787,usec=4544,usec_per_call=5.77,rejected_calls=0,failed_calls=0", ":calls=", ",usec="))
}
