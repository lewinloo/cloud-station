package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
  client *oss.Client
)

var (
  AccessKey = os.Getenv("ALI_AK")
  AccessSecret = os.Getenv("ALI_SK")
  OssEndpoint = os.Getenv("ALI_OSS_ENDPOINT")
  BucketName = os.Getenv("ALI_BUCKET_NAME")
)

func init () {
  c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
  if err != nil {
    panic(err)
  }
  client = c
}

func TestBucketList(t *testing.T) {
  lsRes, err := client.ListBuckets()
  if err != nil {
    t.Log(err)
  }

  for _, bucket := range lsRes.Buckets {
      fmt.Println("Buckets:", bucket.Name)
  }
}

func TestUploadFile(t *testing.T) {
  bucket, err := client.Bucket("my-bucket")
    if err != nil {
      t.Log(err)
    }

    err = bucket.PutObjectFromFile("my-object", "LocalFile")
    if err != nil {
      t.Log(err)
    }
}
