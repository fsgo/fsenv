/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/2
 */

package appenv

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppRoot(t *testing.T) {
	wd, _ := os.Getwd()

	tests := []struct {
		name string
		want string
	}{
		{
			name: "case 1",
			want: filepath.Dir(filepath.Dir(wd)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppRoot()
			if got != tt.want {
				t.Errorf("AppRoot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDirMatch(t *testing.T) {
	type args struct {
		baseDir   string
		fileNames []string
	}
	tests := []struct {
		name    string
		args    args
		wantDir string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDir, err := findDirMatch(tt.args.baseDir, tt.args.fileNames)
			if (err != nil) != tt.wantErr {
				t.Errorf("findDirMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDir != tt.wantDir {
				t.Errorf("findDirMatch() gotDir = %v, want %v", gotDir, tt.wantDir)
			}
		})
	}
}
