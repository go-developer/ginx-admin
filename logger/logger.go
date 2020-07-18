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
	Business *goLogger.WrapperLogger
	// Access 访问日志
	//
	// Author : go_developer@163.com<张德满>
	Access *goLogger.WrapperLogger
)

// InitLogger 初始化日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/29 19:48:25
func InitLogger() {
	var (
		businessLogCfg  *goLogger.LogConfig
		accessLogConfig *goLogger.LogConfig
	)
	businessLogCfg = goLogger.BuildLogConfig(
		"ginx-admin",
		config.Config.LoggerConfig.Business.Debug,
		zapcore.Level(config.Config.LoggerConfig.Business.Level),
		"json",
		config.Config.LoggerConfig.Business.Path,
		nil,
	)
	Business = goLogger.NewWrapperLogger(businessLogCfg)

	accessLogConfig = goLogger.BuildLogConfig(
		"ginx-admin",
		config.Config.LoggerConfig.Access.Debug,
		zapcore.Level(config.Config.LoggerConfig.Access.Level),
		"json",
		config.Config.LoggerConfig.Access.Path,
		nil,
	)
	Access = goLogger.NewWrapperLogger(accessLogConfig)
	Business.Info(nil, "业务日志 & access日志初始化成功", zap.Any("log_config", config.Config.LoggerConfig))
}
