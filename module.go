// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/3

package fsenv

// IModuleEnv 第三方模块管理环境信息的接口定义
type IModuleEnv interface {
	SetEnvOnce(env IAppEnv)
	Env() IAppEnv
}

// ModuleEnv 用于第三方模块管理环境信息
type ModuleEnv struct {
	moduleEnv IAppEnv
}

// Env 获取环境信息
func (m *ModuleEnv) Env() IAppEnv {
	if m.moduleEnv == nil {
		return Default
	}
	return m.moduleEnv
}

// SetEnvOnce 设置自定义的环境信息，只允许设置一次
func (m *ModuleEnv) SetEnvOnce(env IAppEnv) {
	if m.moduleEnv != nil {
		panic("cannot set  ModuleEnv twice")
	}
	m.moduleEnv = env
}

var _ IModuleEnv = (*ModuleEnv)(nil)
