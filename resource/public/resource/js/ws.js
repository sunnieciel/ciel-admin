let isLogin = "{{if.Session.adminInfo}}1{{else}}0{{end}}"
if (isLogin == 1) {
    $(function () {
        let ws = new WebSocket('ws:{{.Config.server.rootIp}}{{.Config.server.address}}/sys/ws')
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
                let d = JSON.parse(data)
                let utterance = new SpeechSynthesisUtterance(d.msg);
                speechSynthesis.speak(utterance);
                console.log(d.msg)
            }
        } catch (e) {
            alert(e.message)
        }
    })
}
