package aliyun_test

import (
	"os"
	"testing"

	"github.com/lewinloo/cloud-station/store"
	"github.com/lewinloo/cloud-station/store/aliyun"
	"github.com/stretchr/testify/assert"
)

var (
	uploader store.Uploader
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

func TestUpload(t *testing.T) {
	should := assert.New(t)
	err := uploader.Upload(BucketName, "test.txt", "store_testsss.go")
	should.Error(err, "open store_testsss.go: no such file or directory")
	// if should.NoError(err) {
	// 	t.Log("upload ok")
	// }
}

func init() {
	ali, err := aliyun.NewAliOssStore(OssEndpoint, AccessKey, AccessSecret)
	if err != nil {
		panic(err)
	}
	uploader = ali
}
