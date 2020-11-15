// Package util controller层的工具集
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 16:17:45
package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/exception"
)

// Response 响应数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 18:13:12
func Response(ctx *gin.Context, e exception.IException, data interface{}, extra map[string]interface{}) {
	if nil == e {
		e = exception.NewSuccess(data)
	}
	result := gin.H{
		"code":           e.GetCode(),
		"message":        e.GetMessage(),
		"trace_id":       ctx.GetString("trace_id"),
		"system_message": e.GetSystemMessage(),
		"data":           e.GetExceptionData(),
	}
	for k, v := range extra {
		result[k] = v
	}
	ctx.JSON(http.StatusOK, result)
}
