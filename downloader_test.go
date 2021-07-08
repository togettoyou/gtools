package gtools

import (
	"fmt"
	"testing"
)

func TestDownloaderTool(t *testing.T) {
	downloader := NewDownloader()
	onWatch := func(current, total int, percentage float64) {
		fmt.Printf("\r当前已下载大小 %f MB, 下载进度：%.2f%%, 总大小 %f MB",
			float64(current)/1024/1024,
			percentage,
			float64(total)/1024/1024,
		)
	}

	ErrExit(downloader.Download(
		"http://www.baidu.com/img/bd_logo1.png",
		"./temp/",
		onWatch,
	))
	fmt.Printf("\n")
	ErrExit(downloader.SetMulti(false).Download(
		"http://www.baidu.com/img/bd_logo1.png",
		"./temp/bd_logo2.png",
		onWatch,
	))
	fmt.Printf("\n")
}
