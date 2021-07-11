// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/3

package fsenv

import (
	"sync"
)

// WithAppEnv 第三方模块管理环境信息的接口定义
type WithAppEnv interface {
	CanSetAppEnv
	HasAppEnv
}

// CanSetAppEnv 允许设置 app 的 env
type CanSetAppEnv interface {
	SetAppEnv(env AppEnv)
}

// HasAppEnv 可以获取 app env
type HasAppEnv interface {
	AppEnv() AppEnv
}

// withAppEnv 用于第三方模块管理环境信息
type withAppEnv struct {
	moduleEnv AppEnv
	mux       sync.RWMutex
}

// AppEnv 获取环境信息
func (m *withAppEnv) AppEnv() AppEnv {
	m.mux.RLock()
	defer m.mux.RUnlock()

	if m.moduleEnv == nil {
		return Default
	}
	return m.moduleEnv
}

// SetAppEnv 设置自定义的环境信息，只允许设置一次
func (m *withAppEnv) SetAppEnv(env AppEnv) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if m.moduleEnv != nil {
		panic("cannot set  withAppEnv twice")
	}
	m.moduleEnv = env
}

var _ WithAppEnv = (*withAppEnv)(nil)
