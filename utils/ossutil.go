package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"log"
	"math/rand"
	"mime/multipart"
	"time"
)

func UploadFile(file *multipart.FileHeader) (string, error) {
	return uploadFileOss(file)
}

func uploadFileOss(file *multipart.FileHeader) (string, error) {
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
