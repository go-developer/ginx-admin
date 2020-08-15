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

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/ginx-manager/auth"
)

// Login 登陆
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/05 20:16:58
func Login(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	var param struct {
		Account  string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&param); nil != err {
		fmt.Println("参数解析失败, 失败原因 : " + err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
			"reason":  err.Error(),
		})
		return
	}
	token, err := auth.NewLogin().Login(ctx, param.Account, param.Password)
	if err != nil {
		fmt.Println("登录失败, 失败原因 : " + err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
			"reason":  err.Error(),
		})
		return
	}
	fmt.Println("登录的token : " + token)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "请求成功",
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
