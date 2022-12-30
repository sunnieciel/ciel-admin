let month_olympic = [31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
let month_normal = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
let month_name = ["1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"];
let holder = $('.days')
let prev = $(".prev")
let next = $('.next')
let ctitle = $(".calendar-title")
let cyear = $(".calendar-year")
let my_date, my_year, my_month, my_day

function setToToday() {
    my_date = new Date()
    my_year = my_date.getUTCFullYear()
    my_month = my_date.getMonth()
    my_day = my_date.getDate()
}

// 获取某年某月第一天时星期几
function dayStart(month, year) {
    let tempDate = new Date(year, month, 1)
    return (tempDate.getDay())
}

// 获取某月总天数 (计算某年是不是闰年)
function daysMonth(month, year) {
    let temp = year % 4
    if (temp === 0) {
        return month_olympic[month]
    } else {
        return month_normal[month]
    }
}

// 返回日期某月的第一天
function getStarMonth(date) {
    let d = new Date()
    d.setFullYear(date.getFullYear())
    d.setMonth(date.getMonth())
    d.setDate(1)
    return d
}

// 生成月份
async function refreshDate() {
    let res = await getMonthWordsLog(my_year, my_month + 1)
    if (res && res.code !== 0) {
        alert(res.msg)
        return
    }
    let m = new Map()
    if (res && res.data) {
        res.data.forEach(i => {
            let d = new Date(i.createdDate)
            m.set(d.getDate(), i.num)
        })
    }
    let str = ''
    // 获取该月总天数
    let totalDay = daysMonth(my_month, my_year)
    // h阿去该月第一天是星期几
    let firstDay = dayStart(my_month, my_year)
    // 为起始日的日期创建空白节点
    let count = 0
    for (let i = 1; i < firstDay; i++) {
        str += `<li></li>`
        count++
    }
    for (var i = 1; i <= totalDay; i++) {
        count++
        let n = m.get(i) || ''
        let myClass
        let numClass = ''
        if ((i < my_day && my_year == my_date.getFullYear() && my_month == my_date.getMonth()) || my_year < my_date.getFullYear() || (my_year == my_date.getFullYear() && my_month < my_date.getMonth())) {
            myClass = "lightgrey "; //当该日期在今天之前时，以浅灰色字体显示
            if (n) {
                numClass = 'finished-num'
                myClass = 'finished-box'
            }
        } else if (i == my_day && my_year == my_date.getFullYear() && my_month == my_date.getMonth()) {
            myClass = "today-box"; //当天日期以绿色背景突出显示
            numClass = 'today-num'
        } else {
            myClass = "darkgrey"; //当该日期在今天之后时，以深灰字体显示
        }
        str += `
<li>
<p class="${myClass}">${i}</p>
<p class="${numClass}">${n}</p>
</li>` //创建日期节点
        holder.html(str)
        ctitle.html(month_name[my_month])
        cyear.html(my_year)
    }
    if (count === 36) {
        $(".calendar").css('padding-bottom', '56px')
    } else {
        $(".calendar").css('padding-bottom', '10px')
    }
}

async function getMonthWordsLog(year, month) {
    let res = await $.get(`/en/wordStudyLogs?year=${year}&month=${month}`)
    return res
}

setToToday()
refreshDate()

$(".prev").on('click', function () {
    my_month--
    if (my_month < 0) {
        my_year--
        my_month = 11
    }
    refreshDate()
})
$(".next").on('click', function () {
    my_month++
    if (my_month > 11) {
        my_year++
        my_month = 0
    }
    let d = new Date()
    d.setFullYear(my_year)
    d.setMonth(my_month)
    let now = new Date()
    if (getStarMonth(d) > getStarMonth(now)) {
        setToToday()
        return
    }
    refreshDate()
})
