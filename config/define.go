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
