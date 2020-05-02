/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/3
 */

package fsenv

// Default 默认的环境信息
var Default AppEnv

func init() {
	initDefault()
}

func initDefault() {
	rootEnv := &rootEnv{}
	Default = &appEnv{
		rootEnv: rootEnv,
		dataEnv: &dataEnv{
			rootEnv: rootEnv,
		},
		logEnv: &logEnv{
			rootEnv: rootEnv,
		},
		confEnv: &confEnv{
			rootEnv: rootEnv,
		},
	}
}

// RootDir 获取应用根目录
func RootDir() string {
	return Default.RootDir()
}

// SetRootDir 设置应用根目录
func SetRootDir(dir string) {
	Default.SetRootDir(dir)
}

// DataRootDir 设置应用数据根目录
func DataRootDir() string {
	return Default.DataRootDir()
}

// SetDataRootDir 获取应用数据根目录
func SetDataRootDir(dir string) {
	Default.SetDataRootDir(dir)
}

// LogRootDir 获取应用日志根目录
func LogRootDir() string {
	return Default.LogRootDir()
}

// SetLogRootDir 设置应用日志根目录
func SetLogRootDir(dir string) {
	Default.SetLogRootDir(dir)
}

// ConfRootPath 获取应用配置根目录
func ConfRootPath() string {
	return Default.ConfRootPath()
}

// SetConfRootPath 设置应用配置根目录
func SetConfRootPath(dir string) {
	Default.SetConfRootPath(dir)
}
