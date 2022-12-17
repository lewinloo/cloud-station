package aliyun

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/lewinloo/cloud-station/store"
)

var (
	// 对象是否实现了接口的约束
	_ store.Uploader = &AliOssStore{}
)

type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint, accessKey, accessSecret has one empty")
	}
	return nil
}

type AliOssStore struct {
	client *oss.Client
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		Endpoint:     os.Getenv("ALI_OSS_ENDPOINT"),
		AccessKey:    os.Getenv("ALI_AK"),
		AccessSecret: os.Getenv("ALI_SK"),
	})
}

func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AccessSecret)
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
