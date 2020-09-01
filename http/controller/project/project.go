// Package project
//
// Description : 创建项目管理
//
// Author : go_developer@163.com<张德满>
package project

import (
	"github.com/gin-gonic/gin"
)

type createProjectForm struct {
	Name        string `json:"name"`        // 项目名称
	Description string `json:"description"` // 项目描述
	Domain      string `json:"domain"`      // 项目域名
	Port        int    `json:"port"`        // 项目端口
}

// create 创建项目结构体
//
// Author : go_developer@163.com<张德满>
type create struct {
}

// create 创建一个新项目
//
// Author : go_developer@163.com<张德满>
func (c *create) Create(ctx *gin.Context) {

}

// Desc : 解析创建项目参数
//
// Author : go_developer@163.com<张德满>
func (c *create) parseForm(ctx *gin.Context) {

}
