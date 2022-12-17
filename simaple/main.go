package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpoint     = ""
	accessKey    = ""
	accessSecret = ""
	bucketName   = "wecloud-station"
	uploadFile   = ""

	help = false
)

func upload(file_path string) error {
	client, err := oss.New(endpoint, accessKey, accessSecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(file_path, file_path)
	if err != nil {
		return err
	}

	downloadUrl, err := bucket.SignURL(file_path, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}

	fmt.Printf("文件下载URL：%s\n", downloadUrl)
	fmt.Printf("请在一天之内下载\n")

	return nil
}

func validate() error {
	if endpoint == "" || accessKey == "" || accessSecret == "" {
		return fmt.Errorf("endpoint, accessKey, accessSecret has one empty")
	}

	if uploadFile == "" {
		return fmt.Errorf("upload file path required")
	}

	return nil
}

func loadParams() {
	flag.BoolVar(&help, "h", false, "打印帮助信息")
	flag.StringVar(&uploadFile, "f", "", "上传文件名称")
	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <upload_file_path>
Options:
`)

	// 打印有哪些参数可以使用
	flag.PrintDefaults()
}

func main() {
	// 参数加载
	loadParams()

	// 参数校验
	if err := validate(); err != nil {
		fmt.Printf("参数校验异常：%s\n", err)
		usage()
		os.Exit(1)
	}

	// 上传文件
	if err := upload(uploadFile); err != nil {
		fmt.Printf("上传文件失败：%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("文件: %s 上传完成\n", uploadFile)
}
