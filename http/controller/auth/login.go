// Package auth ...
//
// File : login.go
//
// Decs : 身份认证api
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/05 20:12:47
package auth

import "github.com/gin-gonic/gin"

type auth struct{}

// Login 登陆
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/05 20:16:58
func (a *auth) Login() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {}
}
