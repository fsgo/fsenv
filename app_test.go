// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/3

package fsenv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAppEnv(t *testing.T) {
	wd, _ := os.Getwd()
	v := Value{}
	env := NewAppEnv(v)

	v.LogDir = "不会修改"

	got := env.Value()
	want := Value{
		RootDir: wd,
		DataDir: filepath.Join(wd, "data"),
		LogDir:  filepath.Join(wd, "log"),
		ConfDir: filepath.Join(wd, "conf"),
		IDC:     "test",
		RunMode: ModeProduct,
	}
	require.Equal(t, want, got)
	require.NotEmpty(t, got.String())
}
