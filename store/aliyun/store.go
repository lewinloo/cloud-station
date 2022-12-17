package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/lewinloo/cloud-station/store"
)

var (
	// 对象是否实现了接口的约束
	_ store.Uploader = &AliOssStore{}
)

type AliOssStore struct {
	client *oss.Client
}

func NewAliOssStore(endpoint, accessKey, accessSecret string) (*AliOssStore, error) {
	c, err := oss.New(endpoint, accessKey, accessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{client: c}, nil
}

func (s *AliOssStore) Upload(bucketName, objectKey, fileName string) error {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(objectKey, fileName)
	if err != nil {
		return err
	}

	downloadUrl, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}

	fmt.Printf("文件下载URL：%s\n", downloadUrl)
	fmt.Printf("请在一天之内下载\n")

	return nil
}
