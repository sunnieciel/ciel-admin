function logout() {
    $.ajax({
        url: '/admin/logout', type: 'get', dataType: 'json', success: function (data) {
            console.log(data)
            if (data.code == 0) {
                window.location.href = '/login';
            }
        }
    });
}


function updatePwd() {
    let old = prompt('Input your old pwd');
    if (!old) {
        return
    }
    let newPwd = prompt('Input your new pwd');
    if (!newPwd) {
        return
    }
    $.post("/admin/updatePwd", {oldPwd: old, newPwd: newPwd}, (res) => {
        if (res.code == 0) {
            noticeOk('Success')
            location.href = '/login'
        } else {
            noticeError(res.msg)
        }
    });
}

// 添加默认选中
$(".sub-nav a").hide()
let current = $("a[href='" + location.pathname + "']");
$(".sub-nav a[data='" + current.attr("data") + "']").show()
$(".nav a[data='" + current.attr("data") + "']").addClass("link-2-active")


// 监听tab 切换
$(function () {
    $(".nav a").hover(function () {
        $(".nav a").removeClass("link-2-active")
        $(this).addClass("link-2-active")
        $(".sub-nav a").hide()
        $(".sub-nav a[data='" + $(this).attr("data") + "']").show()
    })
})


