package gox

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
)

type FileData struct {
	path       string
	name       string
	dir        string
	ext        string
	extNoPoint string
	httpType   string
	size       int64
	humanSize  string
	isExist    bool
}

func FileInfo(path string) (*FileData, error) {
	info := &FileData{
		path: path,
		dir:  filepath.Dir(path),
		ext:  filepath.Ext(path),
	}

	f, err := os.Open(info.path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	info.name = fi.Name()
	info.size = fi.Size()
	info.humanSize = humanize.Bytes(uint64(info.size))
	info.isExist = true

	// 无 . 的后缀
	extSplit := strings.TrimLeft(info.ext,
		".")
	info.extNoPoint = extSplit

	// 文件类型，取前512字节
	buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		return nil, err
	}
	if len(buf) >= 512 {
		buf = buf[:512]
	}

	info.httpType = http.DetectContentType(buf)

	return info, nil
}

func (f *FileData) IsImage() bool {
	return strings.HasPrefix(f.httpType, "image")
}

func (f *FileData) IsExist() bool {
	return f.isExist
}

func (f *FileData) GetPath() string {
	return f.path
}

func (f *FileData) GetName() string {
	return f.name
}

func (f *FileData) GetDir() string {
	return f.dir
}

func (f *FileData) GetExt() string {
	return f.ext
}

func (f *FileData) GetExtNoPoint() string {
	return f.extNoPoint
}

func (f *FileData) GetHTTPType() string {
	return f.httpType
}

func (f *FileData) GetSize() int64 {
	return f.size
}

func (f *FileData) GetHumanSize() string {
	return f.humanSize
}
