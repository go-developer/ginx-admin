// Package project method 相关方法
//
// Description : project...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020-12-19 9:53 下午
package project

import (
	"github.com/gin-gonic/gin"
	"github.com/go-developer/exception"
	"github.com/go-developer/ginx-admin/define"
	"github.com/go-developer/ginx-admin/http/util"
	"github.com/go-developer/ginx-manager/core"
	commonUtil "github.com/go-developer/go-util/util"
)

var (
	// Method 请求方法
	Method *method
)

func init() {
	Method = &method{}
}

type method struct {
}

// GetMethodList 获取请求方法列表
//
// Author : go_developer@163.com<张德满>
func (m *method) GetMethodList(ctx *gin.Context) {
	var (
		pageStr string
		sizeStr string
		page    int
		size    int64
	)

	pageStr = ctx.DefaultQuery("page", "1")
	sizeStr = ctx.DefaultQuery("size", "10")
	commonUtil.ConvertAssign(&page, pageStr)
	commonUtil.ConvertAssign(&size, sizeStr)
	if list, err := core.Method.GetMethodByPage(ctx, page, size); nil != err {
		util.Response(ctx, exception.New(define.CodeGetSchemeListError, err.Error(), nil), nil, nil)
		return
	} else {
		util.Response(ctx, exception.NewSuccess(list.List), nil, map[string]interface{}{"total": list.Total})
		return
	}
}
