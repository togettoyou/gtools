package gtools

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const suffix = ".tar.gz"

var (
	suffixErr = errors.New("file suffix error")
)

// TarFilesDirs 压缩成 tar.gz 文件
// paths 要压缩的目录或文件
// excludePaths 排除压缩的目录或文件
// tarFilePath 压缩后的文件地址
// onWatch 监听压缩状态
func TarFilesDirs(paths, excludePaths []string, tarFilePath string, onWatch func(path string, written int64)) error {
	if !strings.HasSuffix(tarFilePath, suffix) {
		return suffixErr
	}

	err := MkdirAll(tarFilePath)
	if err != nil {
		return err
	}
	file, err := os.Create(tarFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	gz := gzip.NewWriter(file)
	defer gz.Close()

	tw := tar.NewWriter(gz)
	defer tw.Close()

	for _, i := range paths {
		if err := filepath.Walk(i, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			for _, excludePath := range excludePaths {
				if DirIsContainDir(excludePath, path) {
					return nil
				}
			}

			var link string
			if info.Mode()&os.ModeSymlink == os.ModeSymlink {
				if link, err = os.Readlink(path); err != nil {
					return err
				}
			}

			header, err := tar.FileInfoHeader(info, link)
			if err != nil {
				return err
			}

			header.Name = filepath.Join(filepath.Base(i), strings.TrimPrefix(path, string(filepath.Separator)))
			if err = tw.WriteHeader(header); err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			fh, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fh.Close()

			n, err := io.Copy(tw, fh)
			if err != nil {
				return err
			}
			if onWatch != nil {
				onWatch(path, n)
			}

			return nil
		}); err != nil {
			return err
		}
	}

	return nil
}

// UnTarFile 解压 tar.gz 文件
// dst 解压目标目录
// tarFilePath 要解压的 tar.gz 文件
// onWatch 监听解压状态
func UnTarFile(dst, tarFilePath string, onWatch func(path string, written int64)) error {
	if !strings.HasSuffix(tarFilePath, suffix) {
		return suffixErr
	}

	fr, err := os.Open(tarFilePath)
	if err != nil {
		return err
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case hdr == nil:
			continue
		}

		dstFileDir := filepath.Join(dst, hdr.Name)

		switch hdr.Typeflag {
		case tar.TypeDir:
			fi, err := os.Stat(dstFileDir)
			if b := (err == nil || os.IsExist(err)) && fi.IsDir(); !b {
				if err := os.MkdirAll(dstFileDir, os.ModePerm); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			file, err := os.OpenFile(dstFileDir, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}

			n, err := io.Copy(file, tr)
			file.Close()
			if err != nil {
				return err
			}
			if onWatch != nil {
				onWatch(dstFileDir, n)
			}
		}
	}
}
