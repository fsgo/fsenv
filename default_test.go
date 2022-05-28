// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/3

package fsenv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRootDir(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)
	require.Equal(t, wd, RootDir())

	require.NoError(t, os.Setenv(eKeyRoot, "test"))
	defer func() {
		require.NoError(t, os.Unsetenv(eKeyRoot))
		initDefault()
	}()

	require.Equal(t, "test", RootDir())
	SetRootDir("root_dir")
	require.Equal(t, "root_dir", RootDir())
}

func TestLogRootDir(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)
	require.Equal(t, filepath.Join(wd, "log"), LogRootDir())

	require.NoError(t, os.Setenv(eKeyLog, "test"))
	defer func() {
		require.NoError(t, os.Unsetenv(eKeyLog))
		initDefault()
	}()
	require.Equal(t, "test", LogRootDir())

	SetLogRootDir("log_dir")
	require.Equal(t, "log_dir", LogRootDir())
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
			require.Equal(t, tt.want, got)
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
			require.Equal(t, tt.want, got)
		})
	}

	// 恢复初始化的环境
	initDefault()
}

func TestIDC(t *testing.T) {
	require.Equal(t, "test", IDC())
	require.NoError(t, os.Setenv(eKeyIDC, "gz"))
	defer func() {
		require.NoError(t, os.Unsetenv(eKeyIDC))
		initDefault()
	}()
	require.Equal(t, "gz", IDC())

	SetIDC("jx")
	require.Equal(t, "jx", IDC())
}

func TestRunMod(t *testing.T) {
	require.Equal(t, ModeProduct, RunMode())
	require.NoError(t, os.Setenv(eKeyMode, "dev"))
	defer func() {
		require.NoError(t, os.Unsetenv(eKeyMode))
		initDefault()
	}()
	require.Equal(t, Mode("dev"), RunMode())
	SetRunMode(ModeDebug)
	require.Equal(t, ModeDebug, RunMode())
}
