package utils

import (
	"fmt"
	"strconv"
	"time"
)

const (
	Timezone          = "Asia/Shanghai"
	TimeStdFormat     = "2006-01-02 15:04:05"
	TimeStdDateFormat = "2006-01-02"
)

func GenerateLocalZone() *time.Location {
	localZone, err := time.LoadLocation(Timezone)
	if err != nil {
		fmt.Println(err)
	}

	return localZone
}

// 获取当前的时间字符串
func GetCurrentDateTime() string {
	nowTime := time.Now()
	nowTime = nowTime.In(GenerateLocalZone())
	dateTime := nowTime.Format(TimeStdFormat)

	return dateTime
}

//获取每月的几号
func GetMonthDay() int {
	return time.Now().In(GenerateLocalZone()).Day()
}

/*
获取星期几
注意：如果是周日，系统返回的是0，mms系统的规则是：1,2,3,4,5,6,7
即把0重新赋为7
*/
func GetWeekday() string {
	buf := int(time.Now().In(GenerateLocalZone()).Weekday())
	wday := strconv.Itoa(buf)
	if wday == "0" {
		wday = "7"
	}

	return wday
}

//获取本周周几的日期
func GetDateOfWeek(now time.Time, weekDayStr int) string {
	offset := weekDayStr - int(now.Weekday())
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, GenerateLocalZone()).AddDate(0, 0, offset).Format(TimeStdDateFormat)

	return weekStartDate
}
