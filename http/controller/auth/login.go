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
	var param struct {
		Account  string `json:"account"`
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
	ctx.Header("Access-Control-Allow-Origin", "*") //允许访问所有域
	//ctx.Header("Access-Control-Allow-Headers", "Content-Type") //header的类型
	//ctx.Header("content-type", "application/json")             //返回数据格式是json
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "请求成功",
		"data": map[string]interface{}{
			"token": "test-token",
		},
	})
}

func Verify(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "请求成功",
		"data": map[string]interface{}{
			"username":     17710580607,
			"roles":        []int{1, 2},
			"name":         "张德满",
			"avatar":       "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=2073764831,3773405347&fm=26&gp=0.jpg",
			"introduction": "测试账号",
		},
	})
}
