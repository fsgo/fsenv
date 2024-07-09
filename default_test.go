// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/3

package fsenv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fsgo/fst"
)

func TestRootDir(t *testing.T) {
	wd, err := os.Getwd()
	fst.NoError(t, err)
	fst.Equal(t, wd, RootDir())

	fst.NoError(t, os.Setenv(eKeyRoot, "test"))
	defer func() {
		fst.NoError(t, os.Unsetenv(eKeyRoot))
		initDefault()
	}()

	fst.Equal(t, "test", RootDir())
	SetRootDir("root_dir")
	fst.Equal(t, "root_dir", RootDir())
}

func TestLogRootDir(t *testing.T) {
	wd, err := os.Getwd()
	fst.NoError(t, err)
	fst.Equal(t, filepath.Join(wd, "log"), LogRootDir())

	fst.NoError(t, os.Setenv(eKeyLog, "test"))
	defer func() {
		fst.NoError(t, os.Unsetenv(eKeyLog))
		initDefault()
	}()
	fst.Equal(t, "test", LogRootDir())

	SetLogRootDir("log_dir")
	fst.Equal(t, "log_dir", LogRootDir())
}

func TestConfRootPath(t *testing.T) {
	tests := []struct {
		name string
		call func()
		want string
	}{
		{
			name: "case 1",
			call: func() {
				SetRootDir("./internal")
			},
			want: "internal/conf",
		},
		{
			name: "case 2",
			call: func() {
				initDefault()
				SetConfRootDir("xyz/abc")
			},
			want: "xyz/abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.call()
			got := ConfRootDir()
			fst.Equal(t, tt.want, got)
		})
	}

	// 恢复初始化的环境
	initDefault()
}

func TestDataRootPath(t *testing.T) {
	tests := []struct {
		name string
		call func()
		want string
	}{
		{
			name: "case 1",
			call: func() {
				SetRootDir("./internal")
			},
			want: "internal/data",
		},
		{
			name: "case 2",
			call: func() {
				initDefault()
				SetDataRootDir("xyz/abc")
			},
			want: "xyz/abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.call()
			got := DataRootDir()
			fst.Equal(t, tt.want, got)
		})
	}

	// 恢复初始化的环境
	initDefault()
}

func TestIDC(t *testing.T) {
	fst.Equal(t, "test", IDC())
	fst.NoError(t, os.Setenv(eKeyIDC, "gz"))
	defer func() {
		fst.NoError(t, os.Unsetenv(eKeyIDC))
		initDefault()
	}()
	fst.Equal(t, "gz", IDC())

	SetIDC("jx")
	fst.Equal(t, "jx", IDC())
}

func TestRunMod(t *testing.T) {
	fst.Equal(t, ModeProduct, RunMode())
	fst.NoError(t, os.Setenv(eKeyMode, "dev"))
	defer func() {
		fst.NoError(t, os.Unsetenv(eKeyMode))
		initDefault()
	}()
	fst.Equal(t, Mode("dev"), RunMode())
	SetRunMode(ModeDebug)
	fst.Equal(t, ModeDebug, RunMode())
}
