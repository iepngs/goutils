package datetime

import (
    "fmt"
    "strings"
    "time"
)

const GolangBirth = "2006-01-02 15:04:05"

// 获取某天开始或结束的时间戳
func TheDayOfUnixTime(theDay string, start bool) int64 {
    var format string
    if start {
        format = "%s 00:00:00"
    } else {
        format = "%s 23:59:59"
    }
    datetime := fmt.Sprintf(format, theDay)
    unixTs, _ := time.Parse(GolangBirth, datetime)
    return unixTs.Unix()
}

// 判断当日是当年的第几周
func DayOfYearWeek() string {
    t := time.Now()
    yearDay := t.YearDay()
    yearFirstDay := t.AddDate(0, 0, -yearDay+1)
    firstDayInWeek := int(yearFirstDay.Weekday())
    //今年第一周有几天
    firstWeekDays := 1
    if firstDayInWeek != 0 {
        firstWeekDays = 7 - firstDayInWeek + 1
    }
    var week int
    if yearDay <= firstWeekDays {
        week =  1
    } else {
        week = (yearDay-firstWeekDays)/7 + 2
    }
    if week < 10 {
        return fmt.Sprintf("%d0%d", t.Year(), week)
    }
    return fmt.Sprintf("%d%d", t.Year(), week)
}

// 今日开始时间戳
func TodayStartUnix() int64 {
    t := time.Now()
    return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

// 获取当前时间戳
func NowUnixTime() int64 {
    return time.Now().Unix()
}

// 转换mysql查询结果的日期格式
func FormatMysqlDate(change string) string {
    t, _ := time.Parse(time.RFC3339, change)
    return t.In(time.Local).Format("2006-01-02")
}

// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func DateFormat(t time.Time, format string) string {
    res := strings.Replace(format, "MM", t.Format("01"), -1)
    res = strings.Replace(res, "M", t.Format("1"), -1)
    res = strings.Replace(res, "DD", t.Format("02"), -1)
    res = strings.Replace(res, "D", t.Format("2"), -1)
    res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
    res = strings.Replace(res, "YY", t.Format("06"), -1)
    res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
    res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
    res = strings.Replace(res, "hh", t.Format("03"), -1)
    res = strings.Replace(res, "h", t.Format("3"), -1)
    res = strings.Replace(res, "mm", t.Format("04"), -1)
    res = strings.Replace(res, "m", t.Format("4"), -1)
    res = strings.Replace(res, "ss", t.Format("05"), -1)
    res = strings.Replace(res, "s", t.Format("5"), -1)
    return res
}