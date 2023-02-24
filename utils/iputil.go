package utils

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
)

func GetLocation(ip string) string {
	resp, _ := http.Get(fmt.Sprintf("http://whois.pconline.com.cn/ipJson.jsp?ip=%s&json=true", ip))
	all, _ := ioutil.ReadAll(resp.Body)
	bytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(all)
	data := make(map[string]string)
	json.Unmarshal(bytes, &data)
	s := data["addr"]
	return s
}
