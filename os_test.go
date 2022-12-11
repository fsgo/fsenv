// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_osEnvDefault(t *testing.T) {
	key := "fsenv_k1"
	require.NoError(t, os.Unsetenv(key))
	defer require.NoError(t, os.Unsetenv(key))

	require.Equal(t, "v1", osEnvDefault(key, "v1"))
	require.NoError(t, os.Setenv(key, "v2"))
	require.Equal(t, "v2", osEnvDefault(key, "v1"))
}
