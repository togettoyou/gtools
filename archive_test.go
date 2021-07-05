package gtools

import (
	"testing"
)

func TestArchiveTool(t *testing.T) {
	ErrExit(TarFilesDirs(
		[]string{"./"},
		[]string{"temp", ".git", ".gitignore", ".idea"},
		"temp/gtools.tar.gz",
		func(path string, written int64) {
			t.Logf("成功打包 %s ，写入 %d 字节的数据\n", path, written)
		}),
	)
	ErrExit(UnTarFile(
		"temp/gtools",
		"temp/gtools.tar.gz",
		func(path string, written int64) {
			t.Logf("成功解压 %s ，文件 %d 字节\n", path, written)
		}),
	)
}
