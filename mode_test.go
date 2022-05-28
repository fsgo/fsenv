// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppRunMode(t *testing.T) {
	t.Run("init with value", func(t *testing.T) {
		e1 := NewAppRunModeEnv(ModeDebug)
		require.Equal(t, ModeDebug, e1.RunMode())

		e1.SetRunMode(ModeProduct)
		require.Equal(t, ModeProduct, e1.RunMode())

		e1.SetRunMode(ModeDebug)
		require.Equal(t, ModeDebug, e1.RunMode())
	})

	t.Run("init with empty", func(t *testing.T) {
		e1 := NewAppRunModeEnv("")
		require.Equal(t, ModeProduct, e1.RunMode())
		require.NoError(t, os.Setenv(eKeyMode, "demo"))
		require.Equal(t, Mode("demo"), e1.RunMode())
	})
}
