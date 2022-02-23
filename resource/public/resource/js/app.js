// add put and delete type
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

let moment = 0
let add; // 确定是添加还是修改操作

const noticeOk = (msg) => {
    $.notify(msg, {className: 'success', position: 'top center'})
}
const noticeError = (msg) => {
    $.notify(msg, {className: 'error', position: 'top center'})
}
const noticeWarning = (msg) => {
    $.notify(msg, {className: 'warning', position: 'top center'})
}
const noticeInfo = (msg) => {
    $.notify(msg, {className: 'info', position: 'top center'})
}
const showMain = () => {
    $('.part').hide(moment) // all part hide
    $('main').show(moment)
}
const showDetails = (addType, url) => {
    add = addType
    if (add) {
        $('.part').hide()
        document.getElementsByClassName('details')[0].reset()
        $('.details').show()
    } else {
        $.get(url, (res) => {
            if (res) {
                const {data} = res
                for (let i in data) {
                    $(`.details input[name=${i}]`).val(data[i])
                    $(`.details select[name=${i}]`).val(data[i])
                }
            }
            $("main").hide(moment)
            $('.details').show(moment)
        });
    }
}

const handleUpdatePwd = () => {
    let old = prompt('Input your old pwd');
    if (!old) {
        return
    }
    let newPwd = prompt('Input your new pwd');
    if (!newPwd) {
        return
    }
    $.post("/admin/updatePwd", {oldPwd: old, newPwd: newPwd}, (res) => {
        res.code === 0 ? noticeOk('Success') : alert(res.msg)
    });
}

const getFiles = () => {
    let data = new FormData()
    let files = $('#file')[0].files
    for (let i = 0; i, files.length; i++) {
        let item = files[i]
        if (!item) {
            return data
        }
        data.append('file', item)
    }
    return data
}
const handleUploadFile = () => {
    let data = getFiles();
    data.append("group", $('.details input[name="group"]').val())
    $.ajax({
        url: '/file/upload',
        type: 'post',
        cache: false,
        data: data,
        processData: false,
        contentType: false,
    }).done(res => {
        const {code, msg} = res
        if (code === 0) {
            noticeOk(msg)
            location.reload()
        } else {
            noticeError(msg)
        }
    }).fail(res => {
        console.log(res)
    })
}
const onSubmitDetails = (postUrl, putUrl, data) => {
    if (add === 1) {
        $.post(postUrl, data, (res) => {
            if (res.code === 0) {
                noticeOk('Success')
                setTimeout(() => {
                    location.reload()
                }, 1000)
            } else {
                noticeError(res.msg)
            }
        });
    } else {
        $.put(putUrl, data, (res) => {
            if (res.code === 0) {
                noticeOk('Success')
                showMain()
                setTimeout(() => {
                    location.reload()
                }, 1000)
            } else {
                noticeError(res.msg)
            }
        });
    }
}
const onDel = (id, url) => {
    if (confirm("Are you sure?")) {
        $.delete(`${url}?id=${id}`, (res) => {
            if (res.code === 0) {
                $.notify('Success', {className: 'success', position: 'top center'});
                setTimeout(() => {
                    location.reload()
                }, 1000)
            } else {
                noticeError(res.msg)
            }
        });
    }
}