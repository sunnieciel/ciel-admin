jQuery.each(["put", "delete"], function (i, method) {
    jQuery[method] = function (url, data, callback, type) {
        if (jQuery.isFunction(data)) {
            type = type || callback;
            callback = data;
            data = undefined;
        }
        return jQuery.ajax({
            url: url, type: method, dataType: 'json', data: data, success: callback
        });
    };
});

$(function () {
    // 加载完后让body展示
    $("body").show(33)
    // 刷新后保持页面scrollTo原来的位置
    const keyScrollPos = "scroll"
    let scrollY = sessionStorage.getItem(keyScrollPos);
    if (scrollY) window.scrollTo(0, scrollY);
    // 监听窗体滚动
    $(window).scroll(() => sessionStorage.setItem(keyScrollPos, window.scrollY))
    // 监听tab 切换
    $("#nav a").click(function () {
        $("#nav a").removeClass("link-2-active")
        $(this).addClass("link-2-active")
        $("#sub-nav a").hide()
        $("#sub-nav a[data='" + $(this).attr("data") + "']").show()
    })
    // 移动端菜单监听
    $(".top-mobile").click((e) => {
        $(".top-mobile-menu").toggle()
        $(document).one("click", () => {
            $(".top-mobile-menu").hide()
        })
        e.stopPropagation()
    })
})

const setDark = (v) => {
    Cookies.set('dark', v)
    location.reload()
}
const logout = () => {
    $.ajax({
        url: '/admin/logout', type: 'get', dataType: 'json', success: function (data) {
            console.log(data)
            if (data.code === 0) {
                window.location.href = '/login';
            }
        }
    });
}
const updatePwd = () => {
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


