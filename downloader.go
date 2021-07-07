package gtools

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type downloader struct {
	onStart func(contentLength int) io.Writer
	onEnd   func()
}

func NewDownloader(onStart func(contentLength int) io.Writer, onEnd func()) *downloader {
	return &downloader{
		onStart: onStart,
		onEnd:   onEnd,
	}
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

	if resp.Header.Get("Accept-Ranges") == "bytes" {
		// TODO
	}

	return d.singleDownload(url, filename)
}

func (d *downloader) singleDownload(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := MkdirAll(filename); err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	if d.onStart != nil {
		w := d.onStart(int(resp.ContentLength))
		if w != nil {
			_, err = io.Copy(io.MultiWriter(f, w), resp.Body)
		} else {
			_, err = io.Copy(f, resp.Body)
		}
	} else {
		_, err = io.Copy(f, resp.Body)
	}
	if err != nil {
		return err
	}

	if d.onEnd != nil {
		d.onEnd()
	}
	return nil
}
