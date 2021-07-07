package gtools

import (
	"io"
	"testing"
)

func TestDownloaderTool(t *testing.T) {
	downloader := NewDownloader(
		func(contentLength int) io.Writer {
			t.Logf("开始下载, 文件大小 %f MB", (float64(contentLength))/1024/1024)
			return nil
		},
		func() {
			t.Log("下载成功")
		},
	)
	ErrExit(downloader.Download(
		"http://www.baidu.com/img/bd_logo1.png",
		"./temp/bd_logo1.png",
	))
}
