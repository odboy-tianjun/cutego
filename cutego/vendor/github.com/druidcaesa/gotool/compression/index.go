package compression

import (
	"archive/zip"
	"github.com/druidcaesa/gotool/openfile"
	"io"
	"log"
	"os"
	"strings"
)

type ZipUtils struct {
	f openfile.FileUtils
}

// Compress 压缩文件
//files 文件数组 可以是多目录文件
//dest 压缩文件存放地址
func (z ZipUtils) Compress(files []*os.File, dest string) (bool, error) {
	for _, file := range files {
		//防止用户打开os未关闭，这里调用方法前设置结束关闭
		defer file.Close()
	}
	//判断文件是否存在，存在的话删除
	if z.f.Exists(dest) {
		_, err := z.f.RemoveFile(dest)
		if err != nil {
			log.Fatal(err)
			return false, err
		}
	}
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// DeCompress 解压
func (ZipUtils) DeCompress(zipFile, dest string) (bool, error) {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return false, err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return false, err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return false, err
		}
		w, err := os.Create(filename)
		if err != nil {
			return false, err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return false, err
		}
		w.Close()
		rc.Close()
	}
	return true, nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)
	if start < 0 || start > length {
		panic("start is wrong")
	}
	if end < start || end > length {
		panic("end is wrong")
	}
	return string(rs[start:end])
}
