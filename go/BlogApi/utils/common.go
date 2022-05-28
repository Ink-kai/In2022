package utils

import "time"

const (
	YearMonthDay     = "2006-01-02"
	HourMinuteSecond = "15:04:05"
	DefaultLayout    = YearMonthDay + " " + HourMinuteSecond
)

// 默认格式日期字符串转time
func TimeStrToTimeDefault(str string) time.Time {
	parseTime, _ := time.ParseInLocation(DefaultLayout, str, time.Local)
	return parseTime
}

// 指定格式日期字符串转time
func TimeStrToTime(str, layout string) time.Time {
	parseTime, _ := time.ParseInLocation(layout, str, time.Local)
	return parseTime
}
