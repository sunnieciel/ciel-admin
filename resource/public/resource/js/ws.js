try {
    ws.onopen = function () {
        console.log('连接成功')
    }
    ws.onclose = function () {
        if (ws) {
            ws.close()
            ws = null
        }
        console.log('连接关闭')
    }
    ws.onerror = function () {
        if (ws) {
            ws.close()
            ws = null
        }
        console.log('连接关闭')
    }
    ws.onmessage = function ({data}) {
        speak('你有一条新消息')
        updateNotification()
        switch (location.pathname) {
            case "/admin/walletTopUpApplication":
                location.reload()
                break
        }
    }
} catch (e) {
    alert(e.message)
}
setTimeout(function () {
    document.getElementsByTagName("html")[0].click()
}, 1000)

function updateNotification() {
    let $e = $("#notifications")
    let $n = $e.find('b')
    let num = +$n.html()
    if (num === 0) {
        $e.prop('class', 'btn-warning')
    }
    num++
    $n.html(num)
    Cookies.set(UnreadKey, num)
}

let UnreadKey = 'unreadNum'

function initNotification() {
    let num = +Cookies.get(UnreadKey)
    let $e = $("#notifications")
    $e.prop('class', num === 0 ? 'link-3 fs-13' : 'btn-warning')
    $e.find('b').html(num | 0)
}

$(function () {
    initNotification()
    $("#notifications").on('click', async function (e) {
        e.preventDefault()
        let num = +Cookies.get(UnreadKey)
        if (num !== 0) {
            let res = await $.get('/admin/admin/notifications/clear')
            if (res.code !== 0) {
                alert(res.msg)
            }
        }
        Cookies.set(UnreadKey, 0)
        initNotification()
        location.href = $(this).prop('href')
        return false
    })
})

function speak(sentence) {
    // 创建语音合成器
    const synth = window.speechSynthesis;

    // 设置语音
    const voice = synth.getVoices().find(voice => voice.lang === 'zh-CN');

    // 创建语音文本
    const utterance = new SpeechSynthesisUtterance(sentence);
    utterance.voice = voice;

    // 发出语音
    synth.speak(utterance);
}