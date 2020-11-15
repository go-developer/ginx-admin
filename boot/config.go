// Package boot ...
//
// File : config.go
//
// Decs : 引导加载配置文件
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 20:27:52
package boot

import (
	"github.com/go-developer/exception"
	"github.com/go-developer/ginx-admin/config"
	"github.com/go-developer/ginx-admin/define"
	"github.com/go-developer/go-bootstrap"
	godb "github.com/go-developer/gorm-mysql"
)

// NewConfigBootstrap 加载配置文件
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 20:46:44
func NewConfigBootstrap(configPath string, debugMode bool) bootstrap.AbstractBootstrap {
	return &configBootstrap{
		configPath: configPath,
		debugMode:  debugMode,
	}
}

type configBootstrap struct {
	configPath string
	debugMode  bool
}

func (cb *configBootstrap) Start() error {
	config.InitConfig(cb.configPath, cb.debugMode)
	// TODO 迁移走数据库初始化
	godb.InitDatabaseClient(&godb.DBConfig{
		Read: &godb.Mysql{
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Password: "zhangdeman",
			Database: "ginx",
			Charset:  "utf8",
			Loc:      "Asia/shanghai",
		},
		Write: &godb.Mysql{
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Password: "zhangdeman",
			Database: "ginx",
			Loc:      "Asia/shanghai",
			Charset:  "utf8",
		},
		Log: &godb.MysqlLog{
			OpenLog: false,
		},
	})

	// 注册异常的相关信息
	exception.SetDefaultSuccessCode(define.CodeSuccess)
	exception.SetCodeTable(define.CodeTable)
	return nil
}

func (cb *configBootstrap) Stop() error {
	return nil
}

// 配置文件必须最先加载,所以order足够小,保证最先执行
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 20:34:10
func (cb *configBootstrap) Order() int {
	return -10000
}

func (cb *configBootstrap) GetModuleName() string {
	return "CONFIG-BOOTSTRAP"
}
