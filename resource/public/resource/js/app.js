jQuery.each(["put", "delete"], function (i, method) {
    jQuery[method] = function (url, data, callback, type) {
        if (jQuery.isFunction(data)) {
            type = type || callback;
            callback = data;
            data = undefined;
        }
        return jQuery.ajax({
            url: url,
            type: method,
            dataType: 'json',
            data: data,
            success: callback
        });
    };
});
// 监听tab 切换
$(function () {
    $("#nav a").click(function () {
        $("#nav a").removeClass("link-2-active")
        $(this).addClass("link-2-active")
        $("#sub-nav a").hide()
        $("#sub-nav a[data='" + $(this).attr("data") + "']").show()
    })
})

function setDark(value) {
    Cookies.set('dark', value)
    location.reload()
}

function logout() {
    $.ajax({
        url: '/admin/logout', type: 'get', dataType: 'json', success: function (data) {
            console.log(data)
            if (data.code === 0) {
                window.location.href = '/login';
            }
        }
    });
}

function updatePwd() {
    let old = prompt('请输入你的旧密码');
    if (!old) {
        return
    }
    let newPwd = prompt('请输入新密码');
    if (!newPwd) {
        return
    }
    $.put("/admin/updatePwd", {oldPwd: old, newPwd: newPwd}, (res) => {
        if (res.code === 0) {
            alert('success')
            location.href = '/login'
        } else {
            alert(res.msg)
        }
    });
}
