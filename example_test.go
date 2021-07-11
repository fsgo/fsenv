// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/24

package fsenv_test

import (
	"fmt"
	"path/filepath"

	"github.com/fsgo/fsenv"
)

func ExampleRootDir() {
	dir := fsenv.RootDir()
	fmt.Println("root dir=", dir)
}

func ExampleConfRootDir() {
	dir := fsenv.ConfRootDir()
	fmt.Println("conf root dir=", dir)
	// 输出 {xxx}/conf

	dbUserConfPath := filepath.Join(fsenv.ConfRootDir(), "db", "db_user.toml")
	fmt.Println("db user confPath=", dbUserConfPath)
	// 输出 {xxx}/conf/db/db_user.toml
}

func ExampleDataRootDir() {
	dir := fsenv.DataRootDir()
	fmt.Println("data root dir=", dir)
	// 输出 {xxx}/data
}

func ExampleLogRootDir() {
	dir := fsenv.LogRootDir()
	fmt.Println("log root dir=", dir)
	// 输出 {xxx}/log

	dbLogPath := filepath.Join(dir, "db", "mysql.log")
	fmt.Println("db log path=", dbLogPath)
	// 输出 {xxx}/log/db/mysql.log
}
