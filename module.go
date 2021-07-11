// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/3

package fsenv

// WithAppEnv 第三方模块管理环境信息的接口定义
type WithAppEnv interface {
	SetEnv(env AppEnv)
	Env() AppEnv
}

// moduleEnv 用于第三方模块管理环境信息
type moduleEnv struct {
	moduleEnv AppEnv
}

// Env 获取环境信息
func (m *moduleEnv) Env() AppEnv {
	if m.moduleEnv == nil {
		return Default
	}
	return m.moduleEnv
}

// SetEnv 设置自定义的环境信息，只允许设置一次
func (m *moduleEnv) SetEnv(env AppEnv) {
	if m.moduleEnv != nil {
		panic("cannot set  moduleEnv twice")
	}
	m.moduleEnv = env
}

var _ WithAppEnv = (*moduleEnv)(nil)
