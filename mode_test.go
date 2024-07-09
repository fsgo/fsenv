// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/fsgo/fst"
)

func TestAppRunMode(t *testing.T) {
	t.Run("init with value", func(t *testing.T) {
		e1 := NewAppRunModeEnv(ModeDebug)
		fst.Equal(t, ModeDebug, e1.RunMode())

		e1.SetRunMode(ModeProduct)
		fst.Equal(t, ModeProduct, e1.RunMode())

		e1.SetRunMode(ModeDebug)
		fst.Equal(t, ModeDebug, e1.RunMode())
	})

	t.Run("init with empty", func(t *testing.T) {
		e1 := NewAppRunModeEnv("")
		fst.Equal(t, ModeProduct, e1.RunMode())
		fst.NoError(t, os.Setenv(eKeyMode, "demo"))
		fst.Equal(t, Mode("demo"), e1.RunMode())
	})
}
