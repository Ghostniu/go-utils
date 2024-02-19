package timeutils

import (
	"time"
)

const DATE_FORMAT = "2006-01-02"

// 返回某一天的开始时间和结束时间
func GetYeastdayTime(d time.Time) (startTime time.Time, endTime time.Time) {
	startTime = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	endTime = time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
	return
}

// 返回的是这个小时内的开始时间和结束时间
func GetDayTimeByHour(d time.Time) (startTime time.Time, endTime time.Time) {

	startTime = time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), 0, 0, 0, d.Location())
	endTime = time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), 59, 59, 0, d.Location())
	return
}

// 获取上个月的开始时间和结束时间(字符串类型) week为0本周,-1上周，1下周以此类推
func GetLastWeek(week int) (startTime, endTime time.Time) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}
	year, month, day := now.Date()
	thisWeek := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	startTime = thisWeek.AddDate(0, 0, offset+7*week)
	endTime = thisWeek.AddDate(0, 0, offset+6+7*week)
	return
}

// 获取上个月的开始时间和结束时间(字符串类型)
func GetLastMonth() (startTimeStr string, endTimeStr string) {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	startTimeStr = thisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT)
	endTimeStr = thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT)
	return
}

// 获取昨天的开始时间和结束时间(字符串类型)
func GetYeastDay() (startTimeStr string, endTimeStr string) {
	d := time.Now().Add(-time.Hour * 24)
	startTime := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	endTime := time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
	startTimeStr = startTime.Format(DATE_FORMAT)
	endTimeStr = endTime.Format(DATE_FORMAT)
	return
}

// 获取上个月的开始时间和结束时间(字符串类型) 拆分两半 1 取上半个月的值，2 取下半月的值
func GetLastHalfMonth(year int, month time.Month, day int) (startTimeStr, endTimeStr string) {
	var startTime, endTime time.Time
	// 当月第一天
	dayOne := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// 大于0 说明是在后半段
	if (day - 15) > 0 {
		// 取本月的上半月
		startTime = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(year, month, 15, 0, 0, 0, 0, time.Local)
		return startTime.Format(DATE_FORMAT), endTime.Format(DATE_FORMAT)
	} else {
		// 取上月的后半月
		startTime = time.Date(year, month-1, 16, 0, 0, 0, 0, time.Local)
		endTime = dayOne.AddDate(0, 0, -1)
		return startTime.Format(DATE_FORMAT), endTime.Format(DATE_FORMAT)
	}
}

func StrToTime(startTimeStr string, EndTimeStr string) (startTime time.Time, endTime time.Time, err error) {
	startTime, err = time.Parse(DATE_FORMAT, startTimeStr)
	if err != nil {
		return
	}
	endTime, err = time.Parse(DATE_FORMAT, EndTimeStr)
	if err != nil {
		return
	}
	return
}
