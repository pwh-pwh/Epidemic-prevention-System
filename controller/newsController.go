package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
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
