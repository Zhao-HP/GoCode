package util

import (
	"fmt"
	"github.com/pkg/errors"
	"io/fs"
	"os"
	"path/filepath"
)

// GetFilesByRoot 获取指定目录下的所有文件
func GetFilesByRoot(root string) ([]string, error) {

	files := make([]string, 0)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && len(path) > 0 {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func CreateFile(path string) error {
	f, err := os.Create(path)
	if os.IsPermission(err) {
		return errors.Errorf("权限不足, 无法创建[%s]", path)
	}
	if err != nil {
		fmt.Printf("创建文件[%s]失败, [%s]\n", path, err)
		return errors.Errorf("文件[%s]创建失败", path)
	}
	defer f.Close()

	return nil
}

// FolderIsNotExist 判断目录是否不存在
func FolderIsNotExist(path string) bool {
	return !FolderIsExist(path)
}

// FolderIsExist 判断目录是否存在
func FolderIsExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

// FileIsNotExist 判断文件是否不存在
func FileIsNotExist(path string) bool {
	return !FileIsExist(path)
}

// FileIsExist 判断文件是否存在
func FileIsExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}

	return !fileInfo.IsDir()
}

// HasPerm 判断文件是否有权限
func HasPerm(path string, perm os.FileMode) {
	if FileIsNotExist(path) {
		return
	}
}

// DelFile 删除文件
func DelFile(path string) {
	_ = os.Remove(path)
}
