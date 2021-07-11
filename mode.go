// Copyright(C) 2021 github.com/fsgo  All Rights Reserved.
// Author: fsgo
// Date: 2021/7/11

package fsenv

import (
	"log"
	"sync"
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
	return &runModeEnv{
		mode: mode,
	}
}

type runModeEnv struct {
	mode Mode
	mux  sync.RWMutex
}

func (c *runModeEnv) RunMode() Mode {
	c.mux.RLock()
	defer c.mux.RUnlock()

	if c.mode != "" {
		return c.mode
	}
	return ModeProduct
}

func (c *runModeEnv) SetRunMode(mode Mode) {
	c.mux.Lock()
	c.mode = mode
	c.mux.Unlock()
	// 这个比较特殊，运行在运行期间动态调整

	log.Printf("[fsenv] set RunMode=%q\n", mode)
}

var _ AppRunMode = (*runModeEnv)(nil)
