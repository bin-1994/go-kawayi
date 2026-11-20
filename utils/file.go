package utils

import (
	"os"
	"path"
)

// 文件后缀
func GetExt(filename string) string {
	return path.Ext(filename)
}

// 检查文件是否存在
func IsExist(src string) bool {
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 如果不存在则新建文件夹
func IsNotExistMkDir(path string) error {
	if isExist := IsExist(path); isExist {
		return nil
	}
	if err := MkDir(path); err != nil {
		return err
	}
	return nil
}

// 新建文件夹
func MkDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
