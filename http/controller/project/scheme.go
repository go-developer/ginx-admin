// Package project...
//
// Description : project scheme 维护
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020-09-01 10:58 下午
package project

import (
	"net/http"

	"github.com/go-developer/ginx-manager/core"

	"github.com/gin-gonic/gin"
)

var (
	// Scheme 控制器实例
	Scheme *scheme
)

func init() {
	Scheme = &scheme{}
}

func (s *scheme) Create(ctx *gin.Context) {
	var param struct {
		Scheme string `json:"scheme"`
	}
	if err := ctx.BindJSON(&param); nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
		})
		return
	}
	if result, err := core.Scheme.CreateScheme(ctx, param.Scheme); nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": result,
		})
		return
	}
}

type scheme struct {
}
