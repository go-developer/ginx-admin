// Package main ...
//
// File : ginx-admin.go
//
// Decs : ginx 项目管理后台
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 13:41:47
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-developer/ginx-admin/boot"
)

var (
	//配置文件
	configFilePath string
	//开启debug模式
	debugMode bool
	//帮助
	help bool
)

func main() {
	parseParam()
	boot.Start(configFilePath, debugMode)
	defer boot.Stop()
}

func parseParam() {
	flag.BoolVar(&help, "h", false, "帮助文档")
	flag.BoolVar(&debugMode, "debugMode", false, "是否开启debug模式")
	flag.StringVar(&configFilePath, "configFilePath", "", "配置文件目录")
	flag.Usage = usage
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(-1)
	}
}
func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
	Usage: ginx-admin [-h] [-configFilePath configFilePath] [-debugMode]

	Options:
	`)
	flag.PrintDefaults()
}
