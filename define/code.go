// Package define controller层各种定义状态码定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 16:19:08
package define

const (
	// CodeSuccess 请求成功的状态码
	CodeSuccess = 0
	// CodeParamParseError 参数解析失败
	CodeParamParseError = 4000
)

// CodeTable 错误码表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/15 18:07:11
var CodeTable = map[int]string{
	CodeSuccess:         "请求成功",
	CodeParamParseError: "参数解析失败",
}
