package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	myredis "github.com/pwh-pwh/Epidemic-prevention-System/dao/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

func News(ctx *gin.Context) {
	resp, err := http.Get("https://opendata.baidu.com/data/inner?tn=reserved_all_res_tn&dspName=iphone&from_sf=1&dsp=iphone&resource_id=28565&alr=1&query=%E5%9B%BD%E5%86%85%E6%96%B0%E5%9E%8B%E8%82%BA%E7%82%8E%E6%9C%80%E6%96%B0%E5%8A%A8%E6%80%81")
	if err != nil {
		response.Fail(ctx, "最新疫情新闻获取失败")
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Fail(ctx, "最新疫情新闻获取失败")
		return
	}
	r := gjson.GetBytes(data, "Result.0.items_v2.0.aladdin_res.DisplayData.result.items")
	response.Success(ctx, r.Value())
}

func ChinaData(ctx *gin.Context) {
	redisClient := myredis.GetRedisClient()
	var bytes []byte
	if result, _ := redisClient.Exists(myredis.ChinaData).Result(); result == 1 {
		bytes, _ = redisClient.Get(myredis.ChinaData).Bytes()
	} else {
		resp, _ := http.Get("https://c.m.163.com/ug/api/wuhan/app/data/list-total")
		bytes, _ = ioutil.ReadAll(resp.Body)
		redisClient.Set(myredis.ChinaData, string(bytes), 30*60*time.Second)
	}
	//RawMessage 看作是一部分可以暂时忽略的信息，以后可以进一步去解析，但此时不用。所以，我们保留它的原始形式，还是个字节数组即可。
	message := json.RawMessage(bytes)
	response.Success(ctx, message)
}

func GetRiskArea(ctx *gin.Context) {
	resp, err := http.Get("https://m.sm.cn/api/rest?format=json&method=Huoshenshan.riskArea")
	if err != nil {
		response.Fail(ctx, "疫情数据获取失败")
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Fail(ctx, "疫情数据获取失败")
		return
	}
	r := gjson.GetBytes(data, "data.map")
	response.Success(ctx, r.Value())
}

func GetHistory(ctx *gin.Context) {
	redisClient := myredis.GetRedisClient()
	var dataS string
	if result, _ := redisClient.Exists(myredis.HISTORY_DAY_LIST).Result(); result == 1 {
		dataS = redisClient.Get(myredis.HISTORY_DAY_LIST).Val()
	} else {
		resp, _ := http.Get("https://c.m.163.com/ug/api/wuhan/app/data/list-total")
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			response.Fail(ctx, "疫情数据获取失败")
			return
		}
		r := gjson.GetBytes(data, "data.chinaDayList")
		dataS = r.String()
		redisClient.Set(myredis.HISTORY_DAY_LIST, dataS, time.Second*1800)
	}
	response.Success(ctx, dataS)
}

func InfiniteNews(ctx *gin.Context) {
	id := ctx.Query("id")
	var url string
	if id != "" {
		url = "https://m.sm.cn/api/rest?format=json&method=Huoshenshan.feed&type=latest&id=" + id
	} else {
		url = "https://m.sm.cn/api/rest?format=json&method=Huoshenshan.feed&type=latest"
	}
	resp, _ := http.Get(url)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Fail(ctx, "疫情数据获取失败")
		return
	}
	r := gjson.GetBytes(data, "data")
	response.Success(ctx, r.Value())
}

func VaccineTopData(ctx *gin.Context) {
	response.Fail(ctx, "疫苗数据获取失败")
	//接口失效
	/*redisClient := myredis.GetRedisClient()
	var rd any
	if result, _ := redisClient.Exists(myredis.VaccineTopData).Result(); result == 1 {
		bytes, _ := redisClient.Get(myredis.VaccineTopData).Bytes()
		rd = json.RawMessage(bytes)
	} else {
		resp, err := http.Get("https://api.inews.qq.com/newsqa/v1/automation/modules/list?modules=VaccineTopData")
		if err != nil {
			response.Fail(ctx, "最新疫情新闻获取失败")
			return
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			response.Fail(ctx, "最新疫情新闻获取失败")
			return
		}
		r := gjson.GetBytes(data, "data.VaccineTopData")
		rd = r.Value()
		redisClient.Set(myredis.VaccineTopData, r.String(), 1800*time.Second)
	}
	response.Success(ctx, rd)*/
}

func ChinaVaccineTrendData(ctx *gin.Context) {
	response.Fail(ctx, "疫苗数据获取失败")
}

func Rumor(ctx *gin.Context) {
	var url string
	if offset := ctx.Query("offset"); offset != "" {
		url = "https://c.m.163.com/ug/api/wuhan/app/article/list?limit=6&offset=" + offset
	} else {
		url = "https://c.m.163.com/ug/api/wuhan/app/article/list?limit=6"
	}
	resp, _ := http.Get(url)
	data, _ := ioutil.ReadAll(resp.Body)
	result := gjson.GetBytes(data, "data.items")
	response.Success(ctx, result.String())
}
