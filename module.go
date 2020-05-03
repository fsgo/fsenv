/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/3
 */

package fsenv

// IModuleEnv 第三方模块管理环境信息的接口定义
type IModuleEnv interface {
	SetEnvOnce(env AppEnv)
	Env() AppEnv
}

// ModuleEnv 用于第三方模块管理环境信息
type ModuleEnv struct {
	moduleEnv AppEnv
}

// Env 获取环境信息
func (m *ModuleEnv) Env() AppEnv {
	if m.moduleEnv == nil {
		return Default
	}
	return m.moduleEnv
}

// SetEnvOnce 设置自定义的环境信息，只允许设置一次
func (m *ModuleEnv) SetEnvOnce(env AppEnv) {
	if m.moduleEnv != nil {
		panic("cannot set  ModuleEnv twice")
	}
	m.moduleEnv = env
}

var _ IModuleEnv = (*ModuleEnv)(nil)
