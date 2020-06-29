// Package boot ...
//
// File : boot.go
//
// Decs : 引导服务启动
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 21:12:29
package boot

import "github.com/go-developer/go-bootstrap"

// Start 启动服务
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 21:17:27
func Start(configPath string, debugMode bool) {
	//注册配置文件加载
	bootstrap.ModuleBootstrap.Register(NewConfigBootstrap(configPath, debugMode))
	bootstrap.ModuleBootstrap.Register(NewLoggerBootstrap())
	bootstrap.ModuleBootstrap.Bootstrap()
}

// Stop 停止服务
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 21:17:27
func Stop() {
	bootstrap.ModuleBootstrap.Stop()
}
