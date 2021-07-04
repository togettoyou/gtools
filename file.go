package gtools

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FileToBase64 文件转base64
// path 文件路径
func FileToBase64(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// Base64ToFile base64转文件
// base64Data 要转入的base64数据
// path 保存路径
func Base64ToFile(base64Data, path string) error {
	decodeData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}
	if err := MkdirAll(path); err != nil {
		return err
	}
	return ioutil.WriteFile(path, decodeData, 0666)
}

// CreateFile 新建文件
// data 要写入的数据
// path 保存路径
func CreateFile(data, path string) error {
	if err := MkdirAll(path); err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(data), 0666)
}

// MkdirAll 自动根据路径创建文件夹
func MkdirAll(path string) error {
	folder, _ := filepath.Split(path)
	if !PathIsExist(folder) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// PathIsExist 判断文件或目录是否已存在
func PathIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// GetCurrentPath 获取当前绝对路径
func GetCurrentPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}
