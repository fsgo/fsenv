/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/3
 */

package fsenv

// Default (全局)默认的环境信息
var Default IAppEnv

func init() {
	initDefault()
}

func initDefault() {
	rootEnv := &RootEnv{}
	Default = &AppEnv{
		RootEnv: rootEnv,
		DataEnv: &DataEnv{
			rootEnv: rootEnv,
		},
		LogEnv: &LogEnv{
			rootEnv: rootEnv,
		},
		ConfEnv: &ConfEnv{
			rootEnv: rootEnv,
		},
	}
}

// RootDir (全局)获取应用根目录
func RootDir() string {
	return Default.RootDir()
}

// SetRootDir (全局)设置应用根目录
func SetRootDir(dir string) {
	Default.SetRootDir(dir)
}

// DataRootDir (全局)设置应用数据根目录
func DataRootDir() string {
	return Default.DataRootDir()
}

// SetDataRootDir (全局)获取应用数据根目录
func SetDataRootDir(dir string) {
	Default.SetDataRootDir(dir)
}

// LogRootDir (全局)获取应用日志根目录
func LogRootDir() string {
	return Default.LogRootDir()
}

// SetLogRootDir (全局)设置应用日志根目录
func SetLogRootDir(dir string) {
	Default.SetLogRootDir(dir)
}

// ConfRootDir (全局)获取应用配置根目录
func ConfRootDir() string {
	return Default.ConfRootDir()
}

// SetConfRootDir (全局)设置应用配置根目录
func SetConfRootDir(dir string) {
	Default.SetConfRootDir(dir)
}
