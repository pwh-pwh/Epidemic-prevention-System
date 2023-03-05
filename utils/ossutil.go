package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"time"
)

const Ck = "SESSDATA=58b47afb%2C1693490868%2C6b79b%2A32"
const postUrl = "https://api.bilibili.com/x/dynamic/feed/draw/upload_bfs"
const Origin = "https://message.bilibili.com"
const Refer = "https://message.bilibili.com/"
const File = "file_up"
const biz = "im"
const csrf = "2044cba3f8c9e0138bd703572b1391d0"
const build = "0"
const mobi_app = "web"

func UploadFile(file *multipart.FileHeader) (string, error) {
	return UploadFileB2(file)
}

func UploadFileOss(file *multipart.FileHeader) (string, error) {
	client, err := oss.New("oss-cn-guangzhou.aliyuncs.com", common.Oss.AccessKeyID, common.Oss.AccessKeySecret)
	if err != nil {
		log.Printf("new client error:%v", err)
		return "", err
	}
	bucket, err := client.Bucket("go-pj-oss")
	if err != nil {
		log.Printf("get bucket error:%v", err)
		return "", err
	}
	now := time.Now()
	path := "uploadfile/"
	rand.Seed(now.Unix())
	rNum := rand.Intn(1024 * 64)
	rPath := fmt.Sprintf("%s%v/%v/%v/%v%s", path, now.Year(), now.Month().String(), now.Day(), rNum, file.Filename)
	open, _ := file.Open()
	err = bucket.PutObject(rPath, open)
	if err != nil {
		log.Println(err)
	}
	return "https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/" + rPath, nil
}

//TODO 接入图床
func UploadFileB1(file *multipart.FileHeader) {
	httpClient := http.DefaultClient
	body := new(bytes.Buffer)
	postData := multipart.NewWriter(body)
	part1, _ := postData.CreateFormFile("file_up", "file.png")
	open, _ := file.Open()
	defer open.Close()
	io.Copy(part1, open)
	postData.WriteField("biz", "draw")
	postData.WriteField("category", "daily")
	postData.WriteField("csrf", csrf)
	_ = postData.Close()
	request, _ := http.NewRequest("POST", postUrl, body)
	request.Header.Add("Cookie", Ck)
	//request.Header.Add("User-Agent", UserAgent)
	request.Header.Add("Origin", "https://t.bilibili.com")
	request.Header.Add("Referer", "https://t.bilibili.com")
	request.Header.Add("Content-Type", postData.FormDataContentType())
	response, _ := httpClient.Do(request)
	all, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(all))
}

func UploadFileB2(file *multipart.FileHeader) (string, error) {
	httpClient := http.DefaultClient
	body := new(bytes.Buffer)
	postData := multipart.NewWriter(body)
	fileP, _ := postData.CreateFormFile("file_up", file.Filename)
	open, _ := file.Open()
	defer open.Close()
	io.Copy(fileP, open)
	postData.WriteField("biz", biz)
	postData.WriteField("mobi_app", mobi_app)
	postData.WriteField("build", "0")
	postData.WriteField("csrf", csrf)
	postData.Close()
	request, _ := http.NewRequest("POST", postUrl, body)
	request.Header.Add("Cookie", Ck)
	request.Header.Add("Origin", Origin)
	request.Header.Add("Referer", Refer)
	request.Header.Add("Content-Type", postData.FormDataContentType())
	rsp, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}
	all, _ := ioutil.ReadAll(rsp.Body)
	jsonData := make(map[string]any)
	_ = json.Unmarshal(all, &jsonData)
	if v, ok := jsonData["data"].(map[string]any); ok {
		url, ok := v["image_url"]
		urlS := url.(string)
		if ok {
			return urlS, nil
		} else {
			return "", errors.New("解析数据失败")
		}
	} else {
		return "", errors.New("解析数据失败")
	}
}
