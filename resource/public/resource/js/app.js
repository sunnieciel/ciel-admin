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

function noticeOk(msg) {
    $.notify(msg, {className: 'success', position: 'top center'})
}

function noticeOkEle(e, msg) {
    e.notify(msg, {className: 'success', position: 'right'})
}

function noticeError(msg) {
    $.notify(msg, {className: 'error', position: 'top center'})
}

function noticeErrorEle(e, msg) {
    e.notify(msg, {className: 'error', position: 'right'})
}

function noticeWarning(msg) {
    $.notify(msg, {className: 'warning', position: 'top center'})
}

function noticeInfo(msg) {
    $.notify(msg, {className: 'info', position: 'top center'})
}

function logout() {
    $.ajax({
        url: '/admin/logout',
        type: 'get',
        dataType: 'json',
        success: function (data) {
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
        res.code === 0 ? noticeOkEle($(this), 'Success') : noticeErrorEle($(this), res.msg);
    });
}


// 添加默认选中
$("#secondary-tabs a").hide()
let current = $("a[href='" + location.pathname + "']");
$("#secondary-tabs a[data='" + current.attr("data") + "']").show()
$("#tabs a[data='" + current.attr("data") + "']").addClass("tab-current")

// 去除 pagination 里面多余的样式
$(".pagination span").each(function () {
    if (isNaN($(this).text())) $(this).removeClass("GPageSpan")
})
$(function () {
    // 监听tab 切换
    $("#tabs a").click(function () {
        $("#tabs a").removeClass("tab-current")
        $(this).addClass("tab-current")
        $("#secondary-tabs a").hide()
        $("#secondary-tabs a[data='" + $(this).attr("data") + "']").show()
    })
    // 监听搜索 input enter 事件
    $("#search input").keydown(function (e) {
        if (e.keyCode === 13) $("#search").submit()
    })
})

