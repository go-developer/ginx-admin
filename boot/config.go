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
	"github.com/go-developer/ginx-admin/config"
	"github.com/go-developer/go-bootstrap"
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
