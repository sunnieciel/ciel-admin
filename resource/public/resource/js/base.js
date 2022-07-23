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
let msgNum = $(".msg-num") // unread msg element
const keyScrollPos = "scroll" // key scroll
$(async function () {
    // show body after page loaded
    $("body").show(3)
    // listen window scroll event
    $(window).scroll(() => sessionStorage.setItem(keyScrollPos, window.scrollY))
    // handle scroll to original position
    handleScrollToOriginalPosition()
    // handle top link switch event
    $("#nav a").click(handleTopLinkSwitch)
    // handle top mobile menu event
    $(".top-mobile").click(handleMobileMenu)
    // listen msg unread num change
    msgNum.change(listenMsgUnreadNum)
    // clear unread msg count
    msgNum.click(handleUnreadMsg)
    // check unread msg count
    checkUnreadMsgCount()

})

// handle keep original position when window reload
function handleScrollToOriginalPosition() {
    let scrollY = sessionStorage.getItem(keyScrollPos);
    if (scrollY) window.scrollTo(0, scrollY);
}

// handle top link switch event
function handleTopLinkSwitch() {
    $("#nav a").removeClass("link-2-active")
    $(this).addClass("link-2-active")
    $("#sub-nav a").hide()
    $("#sub-nav a[data='" + $(this).attr("data") + "']").show()
}

// switch dark model
const setDark = (v) => {
    Cookies.set('dark', v)
    location.reload()
}
// logout
const logout = () => {
    $.ajax({
        url: '/admin/admin/logout', type: 'get', dataType: 'json', success: function (data) {
            console.log(data)
            if (data.code === 0) {
                window.location.href = '/admin/login';
            }
        }
    });
}
// update password
const updatePwd = () => {
    let old = prompt('请输入你的旧密码');
    if (!old) {
        return
    }
    let newPwd = prompt('请输入新密码');
    if (!newPwd) {
        return
    }
    $.put("/admin/admin/updatePwd", {oldPwd: old, newPwd: newPwd}, (res) => {
        if (res.code === 0) {
            alert('success')
            location.href = '/admin/login'
        } else {
            alert(res.msg)
        }
    });
}
// handle mobile menu show or hide
const handleMobileMenu = (e) => {
    let mobileMenu = $(".top-mobile-menu")
    mobileMenu.toggle()
    $(document).one("click", () => mobileMenu.hide())
    e.stopPropagation()
}

// check unread msg count
const checkUnreadMsgCount = async () => {
    let {data} = await $.get('/admin/adminMessage/unreadMsgCount')
    if (data !== 0) {
        msgNum.attr("data", data).change()
    }
}
// clear unread msg count
const handleUnreadMsg = async () => {
    let {code, msg} = await $.get('/admin/adminMessage/clearUnread')
    if (code != 0) {
        alert(msg)
    }
    msgNum.attr("data", 0).change()
}
// listen unread msg num
const listenMsgUnreadNum = () => {
    if (msgNum.attr("data") == 0) {
        msgNum.hide()
    } else {
        msgNum.text(`您有${+msgNum.attr("data")}条新的消息`)
        msgNum.show()
    }
}


