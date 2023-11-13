package gox

import (
	"reflect"
	"testing"
	"time"
)

func TestFileInfo(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *FileData
		wantErr bool
	}{
		{
			name: "is folder",
			args: args{
				path: "../",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "file not exist",
			args: args{
				path: time.Now().String(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				"./file.go",
			},
			want: &FileData{
				path:       "./file.go",
				name:       "file.go",
				dir:        ".",
				ext:        ".go",
				extNoPoint: "go",
				httpType:   "text/plain; charset=utf-8",
				size:       1670,
				humanSize:  "1.7 kB",
				isExist:    true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileInfo(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileInfo() got = %+v,\n want %+v", got, tt.want)
			}
		})
	}
}
