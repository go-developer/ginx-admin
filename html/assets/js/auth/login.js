/**
 * 登录方法
 */
$("#login_button").click(function () {
    const data = {
        "account":$("#account").val(),
        "password":$("#password").val(),
    }
    request(methodPost, "/admin/user/login", data, "json", contentTypeJson, loginSuccess)
    logger("账号 :", data.account, "密码 :", data.password)
})

/**
 * 登陆成功的回调
 * @param data
 */
function loginSuccess(data) {
    if(data.code === successCode) {
        setToken(data.data.token, true)
        return
    } else {
        alert(data.message)
    }
}
