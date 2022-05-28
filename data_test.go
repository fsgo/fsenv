// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppDataEnv(t *testing.T) {
	t.Run("init with value", func(t *testing.T) {
		e1 := NewAppDataEnv("demo")
		require.Equal(t, "demo", e1.DataRootDir())
		t.Run("twice", func(t *testing.T) {
			defer func() {
				require.NotNil(t, recover())
			}()
			e1.SetDataRootDir("test")
		})
	})

	t.Run("init with empty", func(t *testing.T) {
		e1 := NewAppDataEnv("")
		require.NoError(t, os.Setenv(eKeyData, "test"))
		defer func() {
			require.NoError(t, os.Unsetenv(eKeyData))
		}()
		require.Equal(t, "test", e1.DataRootDir())
		e1.SetDataRootDir("v2")
		require.Equal(t, "v2", e1.DataRootDir())
	})
}
