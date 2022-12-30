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
$(async function () {
    // show body after page loaded
    $("body").show()
    if (location.pathname.startsWith("/admin")) {
        $("#nav a").click(handleTopLinkSwitch)
    }
    $(".top-mobile").click(handleMobileMenu)
})
// 使页面刷新后滚动到原先的位置
// $(function () {
//     const keyScrollPos = "scroll" // key scroll
//     $(window).scroll(() => sessionStorage.setItem(keyScrollPos, window.scrollY))
//     handleScrollToOriginalPosition()
//     function handleScrollToOriginalPosition() {
//         let scrollY = sessionStorage.getItem(keyScrollPos);
//         if (scrollY) window.scrollTo(0, scrollY);
//     }
// })
//

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
// handle mobile menu show or hide
const handleMobileMenu = (e) => {
    let mobileMenu = $(".top-mobile-menu")
    mobileMenu.toggle()
    $(document).one("click", () => mobileMenu.hide())
    e.stopPropagation()
}
// admin  logout
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
// admin update password
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

function goTop() {
    event.preventDefault();
    $("html, body").animate({scrollTop: 0}, 1000);
}
