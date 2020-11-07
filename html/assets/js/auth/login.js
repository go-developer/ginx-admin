/**
 * 登录方法
 */
$("#login_button").click(function () {
    const data = {
        "account": $("#account").val(),
        "password": $("#password").val(),
    }
    request(methodPost, "/admin/user/login", data, "json", contentTypeJson, loginSuccess)
})

/**
 * 登陆成功的回调
 * @param data
 */
function loginSuccess(data) {
    if (data.code === successCode) {
        setToken(data.data.token, true)
        var redirectUrl = window.location.hash.replace("#", "")
        if (undefined === redirectUrl || redirectUrl.length === 0) {
            redirectUrl = defaultIndexPage
        }
        logger("登陆后跳转", redirectUrl, "hash地址", window.location.hash)
        redirect(redirectUrl)
        return
    } else {
        alert(data.message)
    }
}
