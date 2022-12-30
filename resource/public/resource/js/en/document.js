let synth = window.speechSynthesis;
// 所有待练习的单词
let words = []
let currentNum = 0 // 练习单词的当前个数
let IsPractice = false // 是否为练习模式
let ParagraphId = 0 //
let BOX // current Box
class Ciel {
    // 段落进度
    paragraphTotalProgress(level) {
        switch (level) {
            case 1:
                return [dark ? 'var(--gray-002)' : 'var(--gray-002)', '普通', 'tag-info']
            case 2:
                return [dark ? 'var(--green-03)' : 'var(--green-03)', '稀有', 'tag-success']
            case 3:
                return [dark ? 'var(--blue-03)' : 'var(--blue-03)', '传承', 'tag-primary']
            case 4:
                return [dark ? 'var(--yellow-06)' : 'var(--yellow-06)', '唯一', 'tag-warning']
            case 5:
                return [dark ? 'var(--brown-02)' : 'var(--brown-02)', '史诗', 'tag-brown']
            case 6:
                return [dark ? 'var(--purple-02)' : 'var(--purple-02)', '传说', 'tag-purple']
            case 7:
                return [dark ? 'var(--red-03)' : 'var(--red-03)', '神话', 'tag-danger']
        }
    }

    // 单词进度
    wordProgress(level, count) {
        switch (level) {
            case 1:
                return [100 / 33 * count, dark ? 'var(--gray-002)' : 'var(--gray-002)']
            case 2:
                return [100 / 33 * (count - 33), dark ? 'var(--green-03)' : 'var(--green-03)']
            case 3:
                return [100 / 30 * (count - 63), dark ? 'var(--green-03)' : 'var(--green-03)']
            case 4:
                return [100 / 40 * (count - 93), dark ? 'var(--green-03)' : 'var(--green-03)']
            case 5:
                return [100 / 30 * (count - 133), dark ? 'var(--blue-03)' : 'var(--blue-03)']
            case 6:
                return [100 / 30 * (count - 163), dark ? 'var(--blue-03)' : 'var(--blue-03)']
            case 7:
                return [100 / 40 * (count - 193), dark ? 'var(--yellow-06)' : 'var(--yellow-06)']
            case 8:
                return [count - 233, dark ? 'var(--yellow-06)' : 'var(--yellow-06)']
            case 9:
                return [100 / 300 * (count - 333), dark ? 'var(--brown-02)' : 'var(--brown-02)']
            default:
                return [100 / 3333 * (count - 3333), dark ? 'var(--red-03)' : 'var(--red-03)']
        }
    }

    // 单词等级
    wordLevel(level) {
        let cls, desc
        switch (level) {
            case 1:
                cls = "tag-info"
                desc = "单词粉末"
                break
            case 2:
                cls = "tag-success"
                desc = `单词碎片`
                break
            case 3:
                cls = `tag-success`
                desc = `单词结晶体`
                break
            case 4:
                cls = `tag-success`
                desc = `微弱的单词`
                break
            case 5:
                cls = `tag-primary`
                desc = `单词`
                break
            case 6:
                cls = `tag-primary`
                desc = `硕大的单词`
                break
            case 7:
                cls = `tag-warning`
                desc = `闪光的单词`
                break
            case 8:
                cls = `tag-warning`
                desc = `纯洁的单词`
                break
            case 9:
                cls = `tag-brown`
                desc = `燃烧的单词`
                break
            case 10:
                cls = `tag-danger`
                desc = `灿烂的单词`
        }
        return `<a class="${cls} level">${desc}</a>`
    }


    async check() {
        let input = $(".input-practice:focus")
        let data = input.data()
        let errorClass = 'input-practice-error'
        if (data.sentence !== input.val()) {
            if (!input.hasClass(errorClass)) {
                input.addClass(errorClass)
            }
            return
        }
        if (input.hasClass(errorClass)) {
            input.removeClass(errorClass)
        }
        input.val('')
        let res = await $.get(`/en/paragraph/word/addCount?wordId=${data.wordid}`)
        if (res.code !== 0) {
            alert(res.msg)
            return
        }
        if (res.add) {
            let e = $(".today-num")
            e.html(parseInt(e.html()) + 1)
        }
        if (res.data) {
            input.parents('.word-detail').siblings('.level').prop('class', `level ${res.data.class}`).text(res.data.desc)
            data.level = res.data.level
        }
        data.count++
        input.parents('.word-detail').siblings('.num').text(data.count)
        this.speech(data.sentence)
        let progress = this.wordProgress(data.level, data.count)
        input.parents('.word-detail').siblings('.progress-box').find('.progress-current').css({
            width: `${progress[0]}%`, backgroundColor: progress[1]
        })
    }

    // 获取红心♥️
    async getRedHart() {
        let res = await $.get('/user/studyStatus')
        if (res.code !== 0) {
            return alert(res.msg)
        }
        return res.data
    }

    // 获取单词列表
    async getWords(id) {
        let res = await $.get('/en/paragraph/words?paragraphId=' + id)
        if (res.code !== 0) {
            alert(res.msg)
        }
        return res.data
    }

    // 发音
    speech(word, rate) {
        synth.cancel()
        let utterThis = new SpeechSynthesisUtterance(word);
        utterThis.rate = rate ? rate : 0.8
        synth.speak(utterThis)
    }

    // 组合练习的词
    assemblePracticeWords(words) {
        // 隐藏 闯关练习
        let res = ''
        words.forEach((i, index) => {
            let data = this.wordProgress(i.Level.level, i.Level.count);
            let progress = `width: ${data[0]}% background-color:${data[1]}`
            res += `
                <div class="cell flex-center word-box">
                    ${this.wordLevel(i.Level.level)}
                    <span class="color-desc-01">&nbsp;•&nbsp;</span>
                    <span class="link-4 word">${i.Info.en}</span>
                    <span class="color-desc-01">&nbsp;•&nbsp;</span>
                    <div class="progress-box ">
                        <div class="progress-all">
                            <div class="progress-current" style="${progress}"></div>
                        </div>
                    </div>
                    <span class="tag-2 num mr-auto">${i.Level.count}</span>
                    <a class="link-3 mr-6 practice" href="javascript:void(0)" data-count="${index}">练习</a>
                    <div class="word-detail">
                        <b class="word">${i.Info.en}</b>
                        <em class="color-desc-02 fs-10">${i.Info.zh}</em>
                        <p class="sentence"> ${i.Info.sentence} </p>
                        <p class="sentence-zh color-desc-01 fs-10">${i.Info.sentenceZh}</p>
                        <input
                        style="margin-top: 7px"
                                data-sentence="${i.Info.sentence}"
                                data-wordid="${i.Info.id}"
                                data-levelid="${i.Level.id}"
                                data-level="${i.Level.level}"
                                data-count="${i.Level.count}"
                                class="input-practice" type="text" autocomplete="off">
                    </div>
                </div>
                `
        })
        return res
    }

    // 处理单词练习box
    async handleWordPractice(box, check, isPractice) {
        BOX = box
        if (BOX.html().indexOf('闯关失败') !== -1) {
            await makeNewGame()
            return
        }
        /* 赋值到全局变量中，后面提交检查时如果是闯关模式时，如果错了会进行扣小红心 */
        IsPractice = isPractice
        ParagraphId = box.data().id
        // 隐藏单词练习详情
        $('.detail').slideUp(333)
        // 隐藏其他box
        $('.practice-box').each(function (i, ele) {
            if (i !== box.data('index')) {
                // 清空并隐藏
                $(this).empty().slideUp(333)
            }
        })
        // 如果已经有内容则 toggle
        if (check) {
            if (box.html() !== '') {
                /* 如果盒子是关闭的则打开*/
                // 判断小红心是否存在
                let $readHartNum = $(".practice-box .red-hart-num");
                let appear = $readHartNum.css('display') === 'block'
                /*如果是练习模式 并且 小红心是存在的*/
                if (isPractice) {
                    if (appear) { /*是练习模式 但 小红心已经存在了*/
                        // 如果 当前进度不为0 则直接重新获取
                        if (currentNum !== 0) {
                            await makeNewGame()
                        }
                        $readHartNum.hide()
                        if (box.css('display') === 'none') {
                            box.slideDown(333)
                        }
                        return
                    }
                } else {
                    if (!appear) { /*不是练习模式，小红心不存在*/
                        if (currentNum !== 0) {
                            await makeNewGame()
                        }
                        $readHartNum.show()
                        if (box.css('display') === 'none') {
                            box.slideDown(333)
                        }
                        return
                    }
                }
                box.slideToggle()
                return
            }
        }

        // 开始新的游戏
        async function makeNewGame() {
            let res = await $.get(`/en/paragraph/words?paragraphId=${box.data('id')}`)
            words = res.data
            currentNum = 0
            box.html(`
<div class="flex-center">
    <span class="red-hart-num mr-3">♥️ ${heartNum}</span>
    <div class="practice-progress-all mr-3">
        <div class="practice-progress-current"></div>
    </div>
    <span class="fs-10 ml-3 "><span class="current-num">${currentNum}</span>/${words.length}</span>
</div>
<div class="practice-question-title">狗该怎么说?</div>
<div class="practice-options-box"></div>
<div class="practice-notice"></div>`)
            // 出题
            await ciel.genQuestion()
            if (isPractice) {
                $('.practice-box .red-hart-num').hide()
            } else {
                $('.practice-box .red-hart-num').show()
            }
            // 展示
            box.slideDown(333)
        }

        // 获取关键练习单词
        await makeNewGame()
    }

    // 出题
    async genQuestion() {
        // 获取一个没有完成的题
        let question = words.filter(i => !i.ok)
        // 是否完成了？
        if (question.length === 0) {
            $(".practice-options-box").empty()
            $(".practice-question-title").html(`习题已完成 <a class="tag-4 continue" href="javascript:void(0)">继续</a> <a class="tag-4 breakThrough" href="javascript:void(0)">闯关</a>`)
            /* if is practice add read hear num */
            if (IsPractice && +heartNum !== 5) {
                let res = await $.post('/user/addHeart')
                if (res.code !== 0) {
                    alert(res.msg)
                }
                heartNum = res.data.redHardNum
                $('.red-hart-num').html(`♥️ ${res.data.redHardNum}`)
            }
            /*if user is breaking through add paragraph progress */
            if (!IsPractice) {
                let res = await $.post('/en/paragraph/addProgress', {id: ParagraphId})
                let e = BOX.prev().prev().prev()
                let d = ciel.paragraphTotalProgress(res.data.level)
                console.log(d)
                e.children('span:first-child').prop('class', d[2]).html(d[1])
                d = ciel.paragraphTotalProgress(res.data.level)
                e.find('.progress-current').css({
                    width: `${res.data.progress}%`, backgroundColor: d[0]
                })
            }
            return
        }
        question = question[Math.floor(Math.random() * question.length)]
        this.chineseQuestion(question)
    }

    // 生成中文问题
    chineseQuestion(q) {
        // 获取选项
        let options = this.getOptions(q)
        options = Array.from(options).sort(() => .5 - Math.random())
        // 设置问题
        $('.practice-question-title').data(q.Info).html(`"${q.Info.zh}" 怎么说`)
        // 设置选择
        $('.practice-options-box').empty()
        options.forEach(i => {
            let e = $(`<div class="practice-option">${i.Info.en}</div>`)
            e.data(i.Info)
            $('.practice-options-box').append(e)
        })
    }

    // 获取选项
    getOptions(q) {
        let options = new Set()
        options.add(q)
        let len = 4 // 默认4个选项
        if (words.length <= len) {
            len = words.length
        }
        let count = 0
        while (options.size !== len) {
            count++
            options.add(words[Math.floor(Math.random() * words.length)])
            if (count > 10) {
                alert('死循环')
                return
            }
        }
        return options;
    }

    //检查练习单词选择
    async checkPracticeWord(e) {
        //  获取题目
        let question = $(".practice-question-title").data()
        this.speech(e.data().en)
        /* if correct*/
        if (question.en === e.data().en) {
            // 设置对应的问题已完成
            for (let i = words.length - 1; i >= 0; i--) {
                if (words[i].Info.en === question.en) {
                    words[i].ok = 1
                }
            }
            currentNum++
            $(".current-num").html(currentNum)
            $(".practice-progress-current").css('width', `${100 / words.length * currentNum}%`)
            e.removeClass('practice-option').addClass('practice-option-success')
            setTimeout(() => {
                $('.practice-notice').hide()
                this.genQuestion()
            }, 1000)
        } else {
            e.removeClass('practice-option').addClass('practice-option-error')
            $('.practice-notice').html(`
            <b>正确答案</b>
            ${question.en}
            `).show()
            /*if error and the use is breaking through*/
            if (!IsPractice) {
                let res = await $.post('/user/deductionHeart')
                if (res.code !== 0) {
                    alert(res.msg)
                    return
                }
                heartNum = res.data
                $('.red-hart-num').html(`♥️ ${res.data}`)
                if (res.data === 0) {
                    setTimeout(function () {
                        $('.practice-box').html('闯关失败！')
                    }, 1000)
                }
            }
        }
    }
}

const ciel = new Ciel()

$(function () {
    $('h2').each(function (i, e) {
        let h2 = $(this)
        h2.prop('id', `${h2.text()}`)
        $(".cell-nav-box ul").append($(`<li><a class="" href="#${h2.text()}">${h2.text()}</a></li>`))
    })
    $(".progress-current").each(function () {
        $(this).css({
            width: $(this).data().progress + '%', backgroundColor: ciel.paragraphTotalProgress($(this).data().level)[0]
        },)
    })
    // 点击 单词
    $(".word").on('click', async function (e) {
        e.preventDefault()

        // 隐藏 闯关练习
        $('.practice-box').slideUp(333)
        // 获取 单词详情元素
        let details = $(this).parent('.paragraph-tool').next().next();
        // 隐藏其他单词详情
        $('.detail').each(function (i, ele) {
            if (i !== details.data('index')) {
                $(ele).slideUp(333)
            }
        })
        // 如果内容为空则请求
        if (details.html().length === 0) {
            // 为空则从服务器获取数据并渲染
            let data = await ciel.getWords($(this).data().id)
            let html = ciel.assemblePracticeWords(data)
            details.hide()
            details.html(html)
            details.slideDown(333)
        } else {
            details.slideToggle(333)
        }
    })
    // 点击 单词练习
    $(".detail").on('click', '.practice', function () {
        $(".word-detail").each(function (index, e) {
            $(this).slideUp(333)
        })
        let d = $(this).siblings('.word-detail')
        if (d.css('display') === 'none') {
            d.slideDown(333)
        } else {
            d.slideUp(333)
        }
    })
    // 点击 闯关练习
    $(".word-practice").on('click', async function (e) {
        e.preventDefault()
        let box = $(this).parent().next().next().next()
        await ciel.handleWordPractice(box, true, true)
        return false
    })

    // 单词练习提交事件
    $(document).on('keypress', function (e) {
        if (e.which === 13) {
            ciel.check()
        }
    })

    // 发音
    $('.content .detail').on('click', '.word,.sentence', function () {
        ciel.speech($(this).text())
    })

    // 闯关
    $(".break-through").on('click', async function (e) {
        e.preventDefault()
        if ($(this).html() === '闯关') {
            // 检查红心是否足够
            const {persistInLearningDays, wordNum, redHardNum} = await ciel.getRedHart();
            if (redHardNum <= 0) {
                alert('红心不足，快去练习吧')
                return
            }
        }
        let box = $(this).parent().next().next().next()
        await ciel.handleWordPractice(box, true)
        return false
    })

    $(".practice-box")
        // 检查 练习答案选择
        .on('click', '.practice-option', function () {
            ciel.checkPracticeWord($(this))
        })
        // 是否继续 闯关练习
        .on('click', '.continue', async function () {
            let e = $(this).parents('.practice-box');
            await ciel.handleWordPractice(e, false, true)
        })
        // 闯关
        .on('click', '.breakThrough', async function () {
            let e = $(this).parents('.practice-box');
            await ciel.handleWordPractice(e, false)
        })
})
