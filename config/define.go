// Package config ...
//
// File : define.go
//
// Decs :  配置文件结构体定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 22:00:13
package config

// Base 服务基础配置
//
// Author : go_developer@163.com<张德满>
type Base struct {
	Port uint `yml:"port"`
}

// Logger 日志相关配置
//
// Author : go_developer@163.com<张德满>
type Logger struct {
	Access   ItemLogger `yml:"access"`
	Business ItemLogger `yml:"business"`
}

// ItemLogger 日志的数据结构
//
// Author : go_developer@163.com<张德满>
type ItemLogger struct {
	Debug bool   `yml:"debug"`
	Level int    `yml:"level"`
	Path  string `yml:"path"`
}
