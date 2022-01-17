package filex

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type FileContext struct {
	Path string
}

type FileInfo struct {
	Name      string
	Size      int64  // 文件大小
	Ext       string // 拓展名，带 .
	ExtNoSpot string // 拓展名，不带 .
	Type      string // http 头部信息文件类型
	Path      string // 传入的文件路径
	Exist     bool   // 文件是否存在
	Error     error
}

func Path(filePath string) *FileContext {
	return &FileContext{Path: filePath}
}

// Info 获取文件信息
func (fc *FileContext) Info() *FileInfo {
	fileInfo := &FileInfo{}
	// 判断文件是否存在
	isExist, err := fc.IsExist()
	fileInfo.Exist = isExist
	if err != nil {
		fileInfo.addError(err)
		return fileInfo
	}

	file, err := os.Open(fc.Path)
	if err != nil {
		fileInfo.addError(err)
		return fileInfo
	}
	info, err := file.Stat()
	if err != nil {
		fileInfo.addError(err)
		return fileInfo
	}

	// 文件类型
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		fileInfo.addError(err)
		return fileInfo
	}
	if len(buffer) >= 512 {
		buffer = buffer[:512]
	}

	fileInfo.Path = fc.Path
	fileInfo.Name = info.Name()
	fileInfo.Size = info.Size()
	fileInfo.Ext = path.Ext(fc.Path)
	fileInfo.ExtNoSpot = strings.TrimLeft(path.Ext(fc.Path), ".")
	fileInfo.Type = http.DetectContentType(buffer)

	return fileInfo
}

// IsExist 文件是否存在，存在返回 true
func (fc *FileContext) IsExist() (exist bool, err error) {
	fileInfo, err := os.Stat(fc.Path)
	if err == nil && fileInfo != nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

// IsExist 文件是否存在
func (fi *FileInfo) IsExist() bool {
	return fi.Exist
}

// IsImage 是否是图片
func (fi *FileInfo) IsImage() bool {
	return strings.HasPrefix(fi.Type, "image")
}

func (fi *FileInfo) addError(err error) {
	if fi.Error == nil {
		fi.Error = err
	} else if err != nil {
		fi.Error = fmt.Errorf("%v; %w", fi.Error, err)
	}
}

// Compress 文件打包，filePaths 待打包的文件路径，packPath 打包后的文件的路径
func Compress(filePaths []string, packPath string) error {
	var files []*os.File
	for _, filePath := range filePaths {
		open, err1 := os.Open(filePath)
		if err1 != nil {
			return err1
		}
		files = append(files, open)
	}

	create, err2 := os.Create(packPath)
	if err2 != nil {
		return err2
	}
	defer create.Close()
	w := zip.NewWriter(create)
	defer w.Close()
	for _, file := range files {
		err3 := doCompress(file, "", w)
		if err3 != nil {
			return err3
		}
	}
	return nil
}

// 打包操作
func doCompress(file *os.File, prefix string, zw *zip.Writer) error {
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
			err = doCompress(f, prefix, zw)
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
