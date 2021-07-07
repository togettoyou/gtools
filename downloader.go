package gtools

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"sync"
)

type downloader struct {
	// 是否开启并发下载 (若支持)
	multi bool
	// 并发协程数 (若支持)
	concurrency int
	// 断点续传 (若支持)
	resume bool
}

func NewDownloader() *downloader {
	return &downloader{
		multi:       true,
		concurrency: runtime.NumCPU(),
		resume:      true,
	}
}

func (d *downloader) SetMulti(multi bool) *downloader {
	d.multi = multi
	return d
}

func (d *downloader) SetConcurrency(concurrency int) *downloader {
	d.concurrency = concurrency
	return d
}

func (d *downloader) SetResume(resume bool) *downloader {
	d.resume = resume
	return d
}

func (d *downloader) Download(url, filename string) error {
	if filename == "" {
		filename = path.Base(url)
	}

	resp, err := http.Head(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("download fail: %s", resp.Status))
	}

	if err := MkdirAll(filename); err != nil {
		return err
	}

	if d.multi && resp.Header.Get("Accept-Ranges") == "bytes" {
		// 支持分段下载
		return d.multiDownload(url, filename, int(resp.ContentLength))
	}

	return d.singleDownload(url, filename)
}

func (d *downloader) singleDownload(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func (d *downloader) multiDownload(url, filename string, contentLen int) error {

	partSize := contentLen / d.concurrency

	wg := sync.WaitGroup{}
	wg.Add(d.concurrency)

	rangeStart := 0

	for i := 0; i < d.concurrency; i++ {
		go func(i, rangeStart int) {
			defer wg.Done()

			rangeEnd := rangeStart + partSize
			// 最后一部分，总长度不能超过 ContentLength
			if i == d.concurrency-1 {
				rangeEnd = contentLen
			}

			downloaded := 0

			// 断点续传
			if d.resume {
				content, err := ioutil.ReadFile(d.getPartFilename(filename, i))
				if err == nil {
					downloaded = len(content)
				}
			}

			d.downloadPartial(url, filename, rangeStart+downloaded, rangeEnd, i)

		}(i, rangeStart)

		rangeStart += partSize + 1
	}

	wg.Wait()

	return d.merge(filename)
}

func (d *downloader) downloadPartial(url, filename string, rangeStart, rangeEnd, i int) {
	if rangeStart >= rangeEnd {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	flags := os.O_CREATE | os.O_WRONLY
	if d.resume {
		flags |= os.O_APPEND
	}

	partFile, err := os.OpenFile(d.getPartFilename(filename, i), flags, 0666)
	if err != nil {
		return
	}
	defer partFile.Close()

	_, err = io.Copy(partFile, resp.Body)
	if err != nil {
		return
	}
}

func (d *downloader) merge(filename string) error {
	destFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer destFile.Close()

	for i := 0; i < d.concurrency; i++ {
		partFileName := d.getPartFilename(filename, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}
		io.Copy(destFile, partFile)
		partFile.Close()
		os.Remove(partFileName)
	}
	return nil
}

func (d *downloader) getPartFilename(filename string, partNum int) string {
	return fmt.Sprintf("%s-%d", filename, partNum)
}
