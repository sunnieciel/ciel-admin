$(function () {
    // 语音播放
    $('i').click(function () {
        $(this).siblings('audio')[0].play()
    }).mouseover(function () {
        $(this).css({cursor: 'pointer'})
    })
    // 例句翻译
    $('.trans').on('click', function () {
        $(this).parent(".en,.zh").slideToggle().siblings('.zh,.en').slideToggle()
    })
})

$(".word").click(function () {
    $(".wordDetails").fadeIn(33)
    let d = $(this).data();
    $(".enName").text(d.en)
    $(".wordVoice").attr('src', '/upload/' + d.wordvoice)[0].play()
    $(".practiceNum").html(d.practicenum)
    $(".failNum").html(d.failnum)
    let res = ""
    d.zh.split("|").forEach(e => {
        let t = e.split(".")
        res += `<span class="color-desc-02 speech">${t[0]}.</span> <span class="mr-6">${t[1]}</span>`
    })
    $(".zhDesc").html(res)
    $(".sentence").html(d.sentence)
    $(".sentenceVoice").attr("src", '/upload/' + d.sentencevoice)
    $(".sentenceZh").html(d.sentencezh)
})
