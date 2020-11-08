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
	"github.com/go-developer/go-util/util"

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

func (s *scheme) GetAll(ctx *gin.Context) {
	if list, err := core.Scheme.GetAllScheme(ctx); nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"total":   len(list),
			"message": "请求成功",
			"data":    list,
		})
		return
	}
}

func (s *scheme) GetSchemeList(ctx *gin.Context) {
	var (
		pageStr string
		sizeStr string
		page    int
		size    int64
	)

	pageStr = ctx.DefaultQuery("page", "1")
	sizeStr = ctx.DefaultQuery("size", "10")
	util.ConvertAssign(&page, pageStr)
	util.ConvertAssign(&size, sizeStr)
	if list, err := core.Scheme.GetSchemeByPage(ctx, page, size); nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "请求成功",
			"data":    list.List,
			"total":   list.Total,
		})
		return
	}
}

func (s *scheme) SoftDelete(ctx *gin.Context) {
	var param struct {
		SchemeID      uint64 `json:"scheme_id"`
		CurrentStatus uint   `json:"current_status"`
	}
	if err := ctx.BindJSON(&param); nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err := core.Scheme.SoftDelete(ctx, param.SchemeID, param.CurrentStatus); nil != err {
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
