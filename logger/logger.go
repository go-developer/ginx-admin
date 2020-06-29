// Package logger ...
//
// File : logger.go
//
// Decs : 日志操作
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/29 18:51:12
package logger

import (
	"github.com/go-developer/ginx-admin/config"
	goLogger "github.com/go-developer/go-logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Business 业务日志
	//
	// Author : go_developer@163.com<张德满>
	Business *zap.Logger
	// Access 访问日志
	//
	// Author : go_developer@163.com<张德满>
	Access *zap.Logger
)

// InitLogger 初始化日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/29 19:48:25
func InitLogger() {
	var err error
	if Business, err = goLogger.NewDefaultLoggerConfig(
		"ginx-admin",
		config.Config.LoggerConfig.Business.Debug,
		zapcore.Level(config.Config.LoggerConfig.Business.Level),
		"json",
		config.Config.LoggerConfig.Business.Path,
		config.Config.LoggerConfig.Business.Path,
	); nil != err {
		panic(err)
	}
	if Access, err = goLogger.NewDefaultLoggerConfig(
		"ginx-admin",
		config.Config.LoggerConfig.Business.Debug,
		zapcore.Level(config.Config.LoggerConfig.Business.Level),
		"json",
		config.Config.LoggerConfig.Business.Path,
		config.Config.LoggerConfig.Business.Path,
	); nil != err {
		panic(err)
	}
	Business.Info("业务日志 & access日志初始化成功", zap.Any("log_config", config.Config.LoggerConfig))
}
