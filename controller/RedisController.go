package controller

import (
	"bufio"
	"github.com/gin-gonic/gin"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/utils"
	"strings"
)

func GetRedisInfo(ctx *gin.Context) {
	redisClient := myredis.GetRedisClient()
	result := make(map[string]any, 3)
	result["dbSize"] = redisClient.DbSize().Val()
	s := redisClient.Info().String()
	sc := bufio.NewScanner(strings.NewReader(s))
	infoMap := make(map[string]string, 20)
	for sc.Scan() {
		text := sc.Text()
		if strings.Contains(text, ":") {
			split := strings.Split(text, ":")
			infoMap[split[0]] = split[1]
		}
	}
	result["info"] = infoMap
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
			Name:  utils.RemoveStart(split[0], "cmdstat_"),
			Value: utils.SubstringBetween(split[1], "calls=", ",usec="),
		})
	}
	result["commandStats"] = csList
	response.Success(ctx, result)
}

type nameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
