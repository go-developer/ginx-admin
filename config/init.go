// Package config ...
//
// File : init.go
//
// Decs : 解析基础配置文件
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 21:36:12
package config

import (
	"fmt"
	"path"

	"github.com/go-developer/go-util/util"
)

// InitConfig 初始化配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 22:08:32
func InitConfig(configPath string, debugMode bool) {
	var (
		err error
	)
	if len(configPath) == 0 {
		if configPath, err = util.ProjectUtil.GetCurrentPath(); nil != err {
			panic("项目根目录获取失败,失败原因 : " + err.Error())
		}
		configPath = path.Join(configPath, "etc")
	}
	Config = &config{
		configPath: configPath,
		debugMode:  debugMode,
	}
	Config.initLogger()
	Config.initBase()
}

// Config 配置文件管理实例
//
// Author : go_developer@163.com<张德满>
var Config *config

type config struct {
	configPath   string
	debugMode    bool
	BaseConfig   Base
	LoggerConfig Logger
}

// initBase 初始化基础配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/29 15:31:30
func (c *config) initBase() {
	var (
		configPath string
		err        error
	)
	configPath = c.getFullConfigPath("base.yml")
	if err = util.YamlUtil.ParseYamlFile(configPath, &c.BaseConfig); nil != err {
		panic("基础配置文件解析失败, 配置文件路径 : " + configPath + " 异常原因 :" + err.Error())
	}
	fmt.Println("配置文件", c.BaseConfig)
}

// initLogger 初始化日志配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/29 14:25:37
func (c *config) initLogger() {
	var (
		configPath string
		err        error
	)
	configPath = c.getFullConfigPath("logger.yml")
	if err = util.YamlUtil.ParseYamlFile(configPath, &c.LoggerConfig); nil != err {
		panic("日志配置文件解析失败, 配置文件路径 : " + configPath + " 异常原因 :" + err.Error())
	}
	fmt.Println("日志配置文件", c.LoggerConfig)
}

// 获取配置文件的绝对路径
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 21:09:31
func (c *config) getFullConfigPath(relativePath string) string {
	return path.Join(c.configPath, relativePath)
}
