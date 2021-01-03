// Package boot ...
//
// File : logger.go
//
// Decs : 引导初始化日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/29 19:50:32
package boot

import (
	"time"

	"github.com/go-developer/ginx-admin/logger"
	"github.com/go-developer/go-bootstrap"
)

// NewLoggerBootstrap 加载配置文件
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 20:46:44
func NewLoggerBootstrap() bootstrap.AbstractBootstrap {
	return &loggerBootstrap{}
}

type loggerBootstrap struct {
}

func (lb *loggerBootstrap) Start() error {
	logger.InitLogger()
	return nil
}

func (lb *loggerBootstrap) Stop() error {
	//将缓存的日志刷入文件
	logger.Business.GetZapLoggerInstance().Sync()
	logger.Access.GetZapLoggerInstance().Sync()
	time.Sleep(time.Second)
	return nil
}

// 日志在业务上优先加载
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 20:34:10
func (lb *loggerBootstrap) Order() int {
	return 0
}

func (lb *loggerBootstrap) GetModuleName() string {
	return "LOGGER-BOOTSTRAP"
}
