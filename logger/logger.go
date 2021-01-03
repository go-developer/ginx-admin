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
	goLogger "github.com/go-developer/logger"
	"github.com/go-developer/logger/wrapper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Business 业务日志
	//
	// Author : go_developer@163.com<张德满>
	Business *wrapper.GinWrapper
	// Access 访问日志
	//
	// Author : go_developer@163.com<张德满>
	Access *wrapper.GinWrapper
)

// InitLogger 初始化日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/29 19:48:25
func InitLogger() {
	var (
		businessCfg *goLogger.RotateLogConfig
		accessCfg   *goLogger.RotateLogConfig
		err         error
	)
	if businessCfg, err = goLogger.NewRotateLogConfig("./logs", "business.log"); nil != err {
		panic("业务日志初始化失败, 失败原因 : " + err.Error())
	}
	if Business, err = wrapper.NewGinWrapperLogger(
		zapcore.DebugLevel,
		true,
		goLogger.GetEncoder(),
		businessCfg,
		[]string{"trace_id"},
	); nil != err {
		panic("业务日志初始化失败, 失败原因 : " + err.Error())
	}

	if accessCfg, err = goLogger.NewRotateLogConfig("./logs", "access.log"); nil != err {
		Business.Panic("请求日志初始化失败", zap.String("err", err.Error()))
	}
	if Access, err = wrapper.NewGinWrapperLogger(
		zapcore.DebugLevel,
		true,
		goLogger.GetEncoder(),
		accessCfg,
		[]string{"trace_id", "response_data"},
	); nil != err {
		Business.Panic("请求日志初始化失败", zap.String("err", err.Error()))
	}
	Business.Info(
		"日志初始化成功",
		zap.String("business_split_config", goLogger.FormatJson(businessCfg)),
		zap.String("access_split_config", goLogger.FormatJson(accessCfg)),
	)
}
