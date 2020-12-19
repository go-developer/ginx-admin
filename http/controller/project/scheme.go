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

	"github.com/go-developer/exception"
	"github.com/go-developer/ginx-admin/define"
	"github.com/go-developer/ginx-admin/http/util"
	"github.com/go-developer/ginx-manager/core"
	commonUtil "github.com/go-developer/go-util/util"

	"github.com/gin-gonic/gin"
)

var (
	// Scheme 控制器实例
	Scheme *scheme
)

func init() {
	Scheme = &scheme{}
}

type scheme struct {
}

// Create 创建一个新的scheme
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 18:28:01
func (s *scheme) Create(ctx *gin.Context) {
	var param struct {
		Scheme string `json:"scheme"`
	}
	if err := ctx.BindJSON(&param); nil != err {
		util.Response(ctx, exception.New(define.CodeParamParseError, err.Error(), nil), nil, nil)
		return
	}
	if result, err := core.Scheme.CreateScheme(ctx, param.Scheme); nil != err {
		util.Response(ctx, exception.New(define.CodeParamParseError, err.Error(), nil), nil, nil)
		return
	} else {
		util.Response(ctx, exception.NewSuccess(result), result, nil)
		return
	}
}

// GetAll 获取全部scheme列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 18:50:45
func (s *scheme) GetAll(ctx *gin.Context) {
	if list, err := core.Scheme.GetAllScheme(ctx); nil != err {
		util.Response(ctx, exception.New(define.CodeParamParseError, err.Error(), nil), nil, nil)
		return
	} else {
		util.Response(ctx, exception.NewSuccess(list), list, nil)

		return
	}
}

// GetSchemeList 获取scheme列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 18:59:53
func (s *scheme) GetSchemeList(ctx *gin.Context) {
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
	if list, err := core.Scheme.GetSchemeByPage(ctx, page, size); nil != err {
		util.Response(ctx, exception.New(define.CodeGetSchemeListError, err.Error(), nil), nil, nil)
		return
	} else {
		util.Response(ctx, exception.NewSuccess(list.List), nil, map[string]interface{}{"total": list.Total})
		return
	}
}

// SoftDelete 软删除scheme
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 19:10:11
func (s *scheme) SoftDelete(ctx *gin.Context) {
	var param struct {
		SchemeID      uint64 `json:"scheme_id"`
		CurrentStatus uint   `json:"current_status"`
	}
	if err := ctx.BindJSON(&param); nil != err {
		util.Response(ctx, exception.New(define.CodeParamParseError, err.Error(), nil), nil, nil)
		return
	}
	if err := core.Scheme.SoftDelete(ctx, param.SchemeID, param.CurrentStatus); nil != err {
		util.Response(ctx, exception.New(define.CodeSoftDeleteSchemeError, err.Error(), nil), nil, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

// Update 更新scheme
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 18:37:24
func (s *scheme) Update(ctx *gin.Context) {
	var param struct {
		SchemeID uint64 `json:"scheme_id"`
		Scheme   string `json:"scheme"`
	}
	if err := ctx.BindJSON(&param); nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := core.Scheme.UpdateScheme(ctx, param.SchemeID, param.Scheme); nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

// GetDetail 获取scheme详情
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 19:25:20
func (s *scheme) GetDetail(ctx *gin.Context) {}
