// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv_test

import (
	"testing"

	"github.com/fsgo/fst"

	"github.com/fsgo/fsenv"
)

type testWith struct {
	fsenv.WithAppEnv
}

func TestWithAppEnv(t *testing.T) {
	a1 := &testWith{}
	fst.NotNil(t, a1.AppEnv())
	fst.SamePtr(t, fsenv.Default, a1.AppEnv())
	e1 := fsenv.NewAppEnv(fsenv.Value{})
	a1.SetAppEnv(e1)
	fst.SamePtr(t, e1, a1.AppEnv())
	t.Run("twice", func(t *testing.T) {
		defer func() {
			re := recover()
			fst.NotNil(t, re)
		}()
		a1.SetAppEnv(e1)
	})
}
