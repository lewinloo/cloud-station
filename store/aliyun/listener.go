package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/schollz/progressbar/v3"
)

func NewDefaultProgressListener() *ProgressListener {
	return &ProgressListener{}
}

type ProgressListener struct {
	bar *progressbar.ProgressBar
}

func (l *ProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		l.bar = progressbar.DefaultBytes(event.TotalBytes, "文件上传中")
	case oss.TransferDataEvent:
		l.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("文件上传成功\n")
	case oss.TransferFailedEvent:
		fmt.Printf("文件上传失败\n")
	}
}
