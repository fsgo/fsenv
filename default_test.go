package fsenv

import (
	"path/filepath"
	"testing"

	"github.com/fsgo/fst"
)

func TestDefault(t *testing.T) {
	doInit()

	fst.NotEmpty(t, AppName())
	root := RootDir()
	fst.NotEmpty(t, root)
	fst.Equal(t, filepath.Join(root, "conf"), ConfDir())
	fst.Equal(t, filepath.Join(root, "data"), DataDir())
	fst.Equal(t, filepath.Join(root, "log"), LogDir())
	fst.Equal(t, filepath.Join(root, "temp"), TempDir())
	fst.Equal(t, IDCOnline, IDC())
	fst.Equal(t, ModeProduct, RunMode())

	SetAttr("k1", "v1")
	got1, ok1 := Attr("k1")
	fst.Equal(t, "v1", got1)
	fst.True(t, ok1)

	SetConfDir("/user/cfg|abs")
	fst.Equal(t, "/user/cfg", ConfDir())

	SetConfDir("/user/cfg")
	fst.Equal(t, filepath.Join(root, "/user/cfg"), ConfDir())

	SetDataDir("/user/data|abs")
	fst.Equal(t, "/user/data", DataDir())

	SetDataDir("/user/data")
	fst.Equal(t, filepath.Join(root, "/user/data"), DataDir())

	SetTempDir("/temp|abs")
	fst.Equal(t, "/temp", TempDir())

	SetTempDir("/temp")
	fst.Equal(t, filepath.Join(root, "temp"), TempDir())

	SetLogDir("/temp/log|abs")
	fst.Equal(t, "/temp/log", LogDir())

	SetLogDir("/temp/log")
	fst.Equal(t, filepath.Join(root, "temp", "log"), LogDir())
}

func TestMustInitWithAppConfPath(t *testing.T) {
	MustInitWithAppConfPath(filepath.Join(".github", "workflows", "go.yml"))
	fst.Equal(t, ".github", AppName())
}
