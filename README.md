# 应用程序运行环境信息

应用程序运行环境信息

[![Build Status](https://travis-ci.org/fsgo/fsenv.png?branch=master)](https://travis-ci.org/fsgo/fsenv)
[![GoCover](https://gocover.io/_badge/github.com/fsgo/fsenv)](https://gocover.io/github.com/fsgo/fsenv)
[![GoDoc](https://godoc.org/github.com/fsgo/fsenv?status.svg)](https://godoc.org/github.com/fsgo/fsenv)


## 1.主要接口
全局方法、接口：
```go
// AppRootDir 自动推断出应用根目录
// 目前是采用向上查找go.mod文件
// 若查找不到则返回当前目录（pwd）
func AppRootDir() string

// RootDir (全局)获取应用根目录
func RootDir() string

// SetRootDir (全局)设置应用根目录
func SetRootDir(dir string)

// DataRootDir (全局)设置应用数据根目录
func DataRootDir() string 

// SetDataRootDir (全局)获取应用数据根目录
func SetDataRootDir(dir string)

// LogRootDir (全局)获取应用日志根目录
func LogRootDir() string

// SetLogRootDir (全局)设置应用日志根目录
func SetLogRootDir(dir string)

// ConfRootDir (全局)获取应用配置根目录
func ConfRootDir() string

// SetConfRootDir (全局)设置应用配置根目录
func SetConfRootDir(dir string)

// NewAppEnv 创建新的应用环境信息
func NewAppEnv(opt Value) AppEnv

// SetIDC (全局) 设置idc
func SetIDC(idc string) 

// IDC (全局)获取应用的 IDC
func IDC() string 

// RunMode (全局)获取应用的运行模式
func RunMode() Mode 

// SetRunMode (全局)设置应用的运行模式
func SetRunMode(mode Mode)
```
注：为了避免环境信息在运行中修改造成混乱，上述`SetXXX`方法只能调用一次。  


## 2.模块使用自己的环境信息
第三方模块管理自己独特的环境信息，而使用全局的,模块可实现下列接口（或者直接继承使用默认的`ModuleEnv`）：
```go
// WithAppEnv 第三方模块管理环境信息的接口定义
type WithAppEnv interface {
	SetAppEnv(env AppEnv)
	AppEnv() AppEnv
}
```
