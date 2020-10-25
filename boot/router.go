// Package boot ...
//
// File : router.go
//
// Decs : 引导路由监听
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/29 19:50:32
package boot

import (
	"fmt"
	"net/http"

	"github.com/go-developer/go-util/util"

	"github.com/go-developer/ginx-admin/http/controller/project"

	"github.com/facebookarchive/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/go-developer/ginx-admin/config"
	"github.com/go-developer/ginx-admin/http/controller/auth"
	"github.com/go-developer/go-bootstrap"
)

// NewRouterBootstrap 加载配置文件
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 20:46:44
func NewRouterBootstrap() bootstrap.AbstractBootstrap {
	return &routerBootstrap{
		router: gin.Default(),
	}
}

type routerBootstrap struct {
	router *gin.Engine
}

func (rb *routerBootstrap) loadRoute() {
	rb.router.POST("/admin/user/login", auth.Login)
	rb.router.GET("/admin/user/info", auth.Verify)
	rb.router.POST("/admin/project/scheme/create", project.Scheme.Create)
	rb.router.GET("/admin/project/scheme/all", project.Scheme.GetAll)
	rb.router.POST("/admin/project/scheme/delete", project.Scheme.SoftDelete)
	rb.router.POST("/admin/project/scheme/update", project.Scheme.Update)
}

func (rb *routerBootstrap) Start() error {
	rb.loadRoute()
	projectPath, _ := util.ProjectUtil.GetCurrentPath()
	rb.router.StaticFS("/auth", http.Dir(projectPath+"/html/auth"))
	rb.router.StaticFS("/assets", http.Dir(projectPath+"/html/assets"))
	if err := gracehttp.Serve(&http.Server{Addr: fmt.Sprintf(":%d", config.Config.BaseConfig.Port), Handler: rb.router}); nil != err {
		panic(err)
	}
	return nil
}

func (rb *routerBootstrap) Stop() error {
	return nil
}

// 路由最后加载,序号足够大
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 20:34:10
func (rb *routerBootstrap) Order() int {
	return 10000
}

func (rb *routerBootstrap) GetModuleName() string {
	return "ROUTER-BOOTSTRAP"
}
