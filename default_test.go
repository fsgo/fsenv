/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/3
 */

package fsenv

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRootDir(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() error:%v", err.Error())
	}

	tests := []struct {
		name string
		want string
	}{
		{
			name: "case 1",
			want: wd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RootDir(); got != tt.want {
				t.Errorf("RootDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogRootDir(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() error:%v", err.Error())
	}
	tests := []struct {
		name string
		want string
	}{
		{
			name: "case 1",
			want: filepath.Join(wd, "log"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogRootDir(); got != tt.want {
				t.Errorf("LogRootDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfRootPath(t *testing.T) {

	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		call func()
		want string
	}{
		{
			name: "case 1",
			args: args{},
			call: func() {
				SetRootDir("./internal")
			},
			want: "internal/conf",
		},
		{
			name: "case 2",
			args: args{},
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
			if got != tt.want {
				t.Errorf("got=%q, want=%q", got, tt.want)
			}
		})
	}

	// 恢复初始化的环境
	initDefault()
}

func TestDataRootPath(t *testing.T) {

	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		call func()
		want string
	}{
		{
			name: "case 1",
			args: args{},
			call: func() {
				SetRootDir("./internal")
			},
			want: "internal/data",
		},
		{
			name: "case 2",
			args: args{},
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
			if got != tt.want {
				t.Errorf("got=%q, want=%q", got, tt.want)
			}
		})
	}

	// 恢复初始化的环境
	initDefault()
}
