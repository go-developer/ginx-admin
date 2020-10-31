
// 公共js

// 是否是开发模式
var isDevelopment = false
if ("localhost" === document.domain || "127.0.0.1" === document.domain) {
    isDevelopment = true
}

if (isDevelopment) {
    console.log("WARN : 当前运行于开发模式")
}

// 日志相关js
/**
 * 公共打印控制台日志
 */
function logger() {
    if (!isDevelopment) {
        return
    }
    console.log(...logger.arguments)
}

// ajax 请求

// 定义 contentType
const contentTypeJson = "application/json";
const contentTypeForm = "application/x-www-form-urlencoded";

// 定义method
const methodGet = "get";
const methodPost = "post";

let adminServiceDomain = "http://localhost:18080"

if (!isDevelopment) {
    adminServiceDomain = ""
}

/**
 * request 统一发送请求
 */
function request(method, url, data, dataType, contentType, successCallback, failCallback) {
    if (undefined === successCallback) {
        successCallback = defaultSuccessCallback
    }

    if (undefined === failCallback) {
        failCallback = defaultFailCallback
    }

    if (contentType === contentTypeJson) {
        data = JSON.stringify(data)
    }
    $.ajax({
        url: adminServiceDomain + url,
        data: data,
        type: method,
        dataType: dataType,
        contentType: contentType,
        success: successCallback,
        crossDomain: true,
        error: failCallback,
    })
}

function defaultSuccessCallback(data) {
    logger("默认的请求成功回调 => 相应数据 : ", data)
}

function defaultFailCallback(xhr, status, error) {
    logger("请求出现异常 : ", xhr, status, error)
}


// 定义 tokenKey
const tokenKey = "GO_GATEWAY_TOKEN"

// 定义api成功状态码
const successCode = 0

// 定义登录页面
const loginPageURI = "/page/auth/login.html"

// 定义默认首页
const defaultIndexPage = "/page/dashboard/index.html"

// 定义当前页面URI
const currentPageURI = window.location.pathname

// 用户信息
let userInfo = {}

let loginToken = $.cookie(tokenKey)

logger("token信息", loginToken)

if (undefined === loginToken || (currentPageURI !== loginPageURI && loginToken.length === 0)) {
    // 非登录页面，并且没有token, 重定向到 登录页面
    redirect(loginPageURI + "#" + currentPageURI)
}

// 校验token有效性
verifyToken()

/**
 * setToken 设置登录token
 * @param token
 * @param remember
 */
function setToken(token, remember) {
    loginToken = token
    if (remember) {
        //记住密码,保存七天
        $.cookie(tokenKey, token, {
            expires: 7,
            path: '/',
            // domain: adminServiceDomain,
            // secure: true
        })
        return
    }
    // 默认浏览器关闭
    $.cookie(tokenKey, token)
}

/**
 * 页面重定向
 * @param url
 */
function redirect(url) {
    window.location.href = url
}

function verifyToken() {
    request(methodGet, "/admin/user/info", "", "json", contentTypeForm, function (data) {
        userInfo = data.data
    })
}