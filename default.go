// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/3

package fsenv

// Default (全局)默认的环境信息
var Default AppEnvRW

func init() {
	initDefault()
}

func initDefault() {
	envRoot := &rootEnv{}
	Default = &appEnv{
		AppRootEnv: envRoot,
		AppDataEnv: &dataEnv{
			rootEnv: envRoot,
		},
		AppLogEnv: &logEnv{
			rootEnv: envRoot,
		},
		AppConfEnv: &confEnv{
			rootEnv: envRoot,
		},
		AppIDCEnv:  &idcEnv{},
		AppRunMode: &runModeEnv{},
	}
}

// RootDir (全局)获取应用根目录
//
//	若没有设置值，会尝试从环境变量 fsenv.root 读取值
//	默认值 go.mod 所在目录 或者 pwd
func RootDir() string {
	return Default.RootDir()
}

// SetRootDir (全局)设置应用根目录
func SetRootDir(dir string) {
	Default.SetRootDir(dir)
}

// DataRootDir (全局)设置应用数据根目录
//
//	若没有设置值，会尝试从环境变量 fsenv.data 读取值
//	默认值 RootDir()/data
func DataRootDir() string {
	return Default.DataRootDir()
}

// SetDataRootDir (全局)获取应用数据根目录
func SetDataRootDir(dir string) {
	Default.SetDataRootDir(dir)
}

// LogRootDir (全局)获取应用日志根目录
//
//	若没有设置值，会尝试从环境变量 fsenv.root 读取值
//	默认值 RootDir()/log
func LogRootDir() string {
	return Default.LogRootDir()
}

// SetLogRootDir (全局)设置应用日志根目录
func SetLogRootDir(dir string) {
	Default.SetLogRootDir(dir)
}

// ConfRootDir (全局)获取应用配置根目录
//
//	若没有设置值，会尝试从环境变量 fsenv.conf 读取值
//	默认值 RootDir()/conf
func ConfRootDir() string {
	return Default.ConfRootDir()
}

// SetConfRootDir (全局)设置应用配置根目录
func SetConfRootDir(dir string) {
	Default.SetConfRootDir(dir)
}

// SetIDC (全局) 设置idc
func SetIDC(idc string) {
	Default.SetIDC(idc)
}

// IDC (全局)获取应用的 IDC
//
//	若没有设置值，会尝试从环境变量 fsenv.idc 读取值
//	默认值为 test
func IDC() string {
	return Default.IDC()
}

// RunMode (全局)获取应用的运行模式
//
//	若没有设置值，会尝试从环境变量 fsenv.mode 读取值
//	默认值为 product
func RunMode() Mode {
	return Default.RunMode()
}

// SetRunMode (全局)设置应用的运行模式
func SetRunMode(mode Mode) {
	Default.SetRunMode(mode)
}
