// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"strings"
)

const (
	eKeyIDC  = "FSENV_IDC"
	eKeyConf = "FSENV_CONF"
	eKeyRoot = "FSENV_ROOT"
	eKeyData = "FSENV_DATA"
	eKeyTemp = "FSENV_TEMP"
	eKeyLog  = "FSENV_LOG"
	eKeyMode = "FSENV_MODE"
)

func osEnvDefault(key string, def string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return def
	}
	return val
}

// parserDirName 解析目录名称，若以 |abs 结尾则表示是绝对路径
func parserDirName(path string) (string, bool) {
	before, after, found := strings.Cut(path, "|")
	if found && after == "abs" {
		return before, true
	}
	return path, false
}
