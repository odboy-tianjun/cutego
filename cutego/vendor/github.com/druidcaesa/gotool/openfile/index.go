package openfile

import (
	"io/ioutil"
	"log"
	"os"
)

type FileUtils struct {
}

// Exists 判断所给路径文件或文件夹是否存在
func (FileUtils) Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// Create 创建文件
func Create(path string) {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		log.Fatalf("db connect error: %#v\n", err.Error())
	}
}

// IsDir 判断所给路径是否为文件夹
func (FileUtils) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func (f FileUtils) IsFile(path string) bool {
	return !f.IsDir(path)
}

// RemoveFile 删除文件，参数文件路径
func (f FileUtils) RemoveFile(path string) (bool, error) {
	//删除文件
	err := os.Remove(path)
	if err != nil {
		return false, err
	}
	return true, nil
}

// OpenFileRdonly 打开文件只读模式 文件不存会直接进项创建
func (f FileUtils) OpenFileRdonly(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
}

// OpenFileWronly 打开文件只写模式，若文件不存在进行创建
func (f FileUtils) OpenFileWronly(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
}

// OpenFileAppend 打开文件在文件后面追加数据，文件文件不存在会进行创建
func (f FileUtils) OpenFileAppend(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
}

// FileCopy 进行文件复制操作
func (f FileUtils) FileCopy(oldPath, newPath string) (bool, error) {
	data, err := ioutil.ReadFile(oldPath)
	if err != nil {
		return false, err
	}
	err = ioutil.WriteFile(newPath, data, 0666)
	if err != nil {
		return false, err
	}
	return true, nil
}
