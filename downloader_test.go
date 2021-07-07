package gtools

import (
	"testing"
)

func TestDownloaderTool(t *testing.T) {
	downloader := NewDownloader()
	ErrExit(downloader.Download(
		"http://www.baidu.com/img/bd_logo1.png",
		"./temp/bd_logo1.png",
	))
}
