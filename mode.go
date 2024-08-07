// Copyright(C) 2020 github.com/fsgo  All Rights Reserved.
// Author: hidu
// Date: 2024/08/07

package fsenv

import "fmt"

// Mode 运行模式
type Mode int32

const (
	// ModeProduct 运行模式-线上发布
	ModeProduct Mode = iota

	// ModeDebug 运行模式-调试
	ModeDebug
)

func (m Mode) String() string {
	switch m {
	case ModeProduct:
		return "product"
	case ModeDebug:
		return "debug"
	default:
		return fmt.Sprintf("unknown(%d)", m)
	}
}

func modeFromEnv(def Mode) Mode {
	switch osEnvDefault(eKeyMode, "") {
	case "product":
		return ModeProduct
	case "debug":
		return ModeDebug
	default:
		return def
	}
}
