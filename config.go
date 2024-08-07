package fsenv

import (
	"os"
	"path/filepath"
)

type Config struct {
	AppName string
	RootDir string
	ConfDir string
	DataDir string
	LogDir  string
	TempDir string
	Attrs   map[string]any
}

func (cfg *Config) MustAttribute() *Attribute {
	attr, err := cfg.Attribute()
	if err != nil {
		panic(err)
	}
	return attr
}

func (cfg *Config) Attribute() (*Attribute, error) {
	root, err := cfg.getRootDir()
	if err != nil {
		return nil, err
	}
	name := cfg.AppName
	if name == "" {
		name = filepath.Base(root)
	}
	attr := NewAttribute(name, root)
	if cfg.ConfDir != "" {
		attr.SetConfDir(cfg.ConfDir)
	}
	if cfg.DataDir != "" {
		attr.SetDataDir(cfg.DataDir)
	}
	if cfg.TempDir != "" {
		attr.SetTempDir(cfg.TempDir)
	}
	if cfg.LogDir != "" {
		attr.SetLogDir(cfg.LogDir)
	}
	for k, v := range cfg.Attrs {
		attr.SetAttr(k, v)
	}
	return attr, nil
}

func (cfg *Config) getRootDir() (string, error) {
	if cfg.RootDir != "" {
		return filepath.Abs(cfg.RootDir)
	}
	return os.Getwd()
}
