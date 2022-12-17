package cli

import (
	"fmt"
	"github.com/lewinloo/cloud-station/store"
	"github.com/lewinloo/cloud-station/store/aliyun"
	"github.com/lewinloo/cloud-station/store/aws"
	"github.com/lewinloo/cloud-station/store/tencent"
	"github.com/spf13/cobra"
)

var (
	ossProvider  string
	ossEndpoint  string
	accessKey    string
	accessSecret string
	bucketName   string
	uploadFile   string
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 上传文件到云商",
	Short:   "upload 上传文件到云商",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)

		switch ossProvider {
		case "aliyun":
			uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
				Endpoint:     ossEndpoint,
				AccessKey:    accessKey,
				AccessSecret: accessSecret,
			})
		case "tencent":
			uploader = tencent.NewTencentOssStore()
		case "aws":
			uploader = aws.NewAwsOssStore()
		default:
			return fmt.Errorf("not support oss storage provider")
		}
		if err != nil {
			return err
		}

		// 使用 Upload 上传文件
		uploader.Upload(bucketName, uploadFile, uploadFile)

		return nil
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider[aliyun/tencent/aws]")
	f.StringVarP(&ossEndpoint, "endpoint", "e", "oss-cn-guangzhou.aliyuncs.com", "oss storage provider endpoint")
	f.StringVarP(&accessKey, "access_key", "k", "", "oss storage provider ak")
	f.StringVarP(&accessSecret, "access_secret", "s", "", "oss storage provider sk")
	f.StringVarP(&bucketName, "bucket_name", "b", "wecloud-station", "oss storage provider bucket name")
	f.StringVarP(&uploadFile, "upload_file", "f", "", "upload file name")
	RootCmd.AddCommand(UploadCmd)
}
