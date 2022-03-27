// var url = "{{.Config.server.adminWs}}";
// var ws = new WebSocket(`${url}?aid={{.Session.adminInfo.Admin.id}}`);
// try {
//     // ws连接成功
//     ws.onopen = function () {
//         console.log("WebSocket Server [" + url + "] 连接成功！");
//     };
//     // ws连接关闭
//     ws.onclose = function () {
//         if (ws) {
//             ws.close();
//             ws = null;
//         }
//         console.log("WebSocket Server [" + url + "] 连接关闭！");
//     };
//     // ws连接错误
//     ws.onerror = function () {
//         if (ws) {
//             ws.close();
//             ws = null;
//         }
//         console.log("WebSocket Server [" + url + "] 连接关闭！");
//     };
//     // ws数据返回处理
//     ws.onmessage = function (result) {
//         const {msg} = JSON.parse(result.data);
//         switch (location.pathname) {
//             case "/topUp/list":
//             case "/withdrawal/list":
//                 window.location.reload();
//                 break
//         }
//         let utterance = new SpeechSynthesisUtterance(msg);
//         speechSynthesis.speak(utterance);
//         noticeInfo(msg)
//     };
// } catch (e) {
//     alert(e.message);
// }