// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/27

package fsenv

import (
	"fmt"
	"os"
	"strings"
)

// LoadEnvFile 加载一个 env 文件到 os.env 里
func LoadEnvFile(fn string) error {
	vs, err := ParserEnvFile(fn)
	if err != nil {
		return err
	}
	for k, v := range vs {
		if err = os.Setenv(k, v); err != nil {
			return fmt.Errorf("s.Setenv(%q,%q) failed: %w", k, v, err)
		}
	}
	return nil
}

// MustLoadEnvFile 加载一个 env 文件到 os.env 里。
// 若有异常，会 panic
func MustLoadEnvFile(fn string) {
	if err := LoadEnvFile(fn); err != nil {
		panic("LoadEnvFile(" + fn + ") failed:" + err.Error())
	}
}

// ParserEnvFile 解析一个 env 文件
func ParserEnvFile(fn string) (map[string]string, error) {
	content, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return ParserEnvContent(content)
}

// ParserEnvContent 解析 env 文件的内容
func ParserEnvContent(content []byte) (map[string]string, error) {
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	result := make(map[string]string, len(lines))
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}
		k, v, err := parserEnvLine(line)
		if err != nil {
			return nil, err
		}
		if len(k) == 0 {
			continue
		}
		result[k] = v
	}
	return result, nil
}

func parserEnvLine(line string) (key string, value string, err error) {
	if strings.HasPrefix(line, "#") {
		return "", "", nil
	}
	arr := strings.SplitN(line, "=", 2)
	if len(arr) != 2 {
		return "", "", fmt.Errorf("invalid line: %q", line)
	}
	key = strings.TrimSpace(arr[0])
	value = strings.TrimSpace(arr[1])
	return key, parserEnvValue(value), nil
}

func parserEnvValue(value string) string {
	if len(value) < 2 {
		return value
	}
	for _, s := range []string{`"`, `'`, "`"} {
		if strings.HasPrefix(value, s) && strings.HasSuffix(value, s) {
			return value[1 : len(value)-1]
		}
	}
	return value
}
