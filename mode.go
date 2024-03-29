// Copyright(C) 2021 github.com/fsgo  All Rights Reserved.
// Author: fsgo
// Date: 2021/7/11

package fsenv

import (
	"os"
	"sync/atomic"
)

// Mode 运行模式
type Mode string

const (
	// ModeDebug 运行模式-调试
	ModeDebug Mode = "debug"

	// ModeProduct 运行模式-线上发布
	ModeProduct Mode = "product"
)

// AppRunMode 机房信息
type AppRunMode interface {
	HasRunMode
	CanSetRunMode
}

// HasRunMode 可以获取 runMode
type HasRunMode interface {
	RunMode() Mode
}

// CanSetRunMode 允许设置 runMode
type CanSetRunMode interface {
	SetRunMode(mode Mode)
}

// NewAppRunModeEnv 创建新的配置环境信息
func NewAppRunModeEnv(mode Mode) AppRunMode {
	ae := &runModeEnv{}
	ae.SetRunMode(mode)
	return ae
}

type runModeEnv struct {
	mode atomic.Value
}

func (c *runModeEnv) RunMode() Mode {
	vs, _ := c.mode.Load().(Mode)
	if len(vs) > 0 {
		return vs
	}
	if val := os.Getenv(eKeyMode); len(val) > 0 {
		return Mode(val)
	}
	return ModeProduct
}

func (c *runModeEnv) SetRunMode(mode Mode) {
	c.mode.Store(mode)
}

var _ AppRunMode = (*runModeEnv)(nil)
