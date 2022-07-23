let id = 0;
$(async function () {
    await getCaptcha()
    $("input").keyup(function (e) {
        if (e.keyCode === 13) {
            handleLogin()
        }
    })

    let res = await $.get('/v1/dict/key/music-url')
    if (res.code === 0) {
        if (res.data !== '') {
            url = res.data
        }
        document.getElementById('music').src = url
    }
})

const getCaptcha = () => {
    id = Math.random();
    $("#captcha").val("")
    $.get("/admin/admin/getCaptcha?id=" + id, function (res) {
        $("#code").attr("src", res.data);
    })
}
