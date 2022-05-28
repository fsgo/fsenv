// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fsgo/fsenv"
)

type testWith struct {
	fsenv.WithAppEnv
}

func TestWithAppEnv(t *testing.T) {
	a1 := &testWith{}
	require.NotNil(t, a1.AppEnv())
	require.Same(t, fsenv.Default, a1.AppEnv())
	e1 := fsenv.NewAppEnv(fsenv.Value{})
	a1.SetAppEnv(e1)
	require.Same(t, e1, a1.AppEnv())
	t.Run("twice", func(t *testing.T) {
		defer func() {
			re := recover()
			require.NotNil(t, re)
		}()
		a1.SetAppEnv(e1)
	})
}
