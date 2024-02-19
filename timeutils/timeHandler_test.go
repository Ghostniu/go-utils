package timeutils

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDayTimeByHour(t *testing.T) {
	d := time.Now()
	yesStart, _ := GetYeastdayTime(d)
	for i := 0; i < 24; i++ {
		d := yesStart.Add(time.Hour * time.Duration(i))
		startTime, endTime := GetDayTimeByHour(d)
		fmt.Println(startTime)
		fmt.Println(endTime)
		fmt.Println()
	}
}

func TestGetLastMonth(t *testing.T) {
	startTimeStr, endTimeStr := GetLastMonth()
	fmt.Println(startTimeStr, endTimeStr)
}

func TestGetLastHalfMonth1(t *testing.T) {
	type args struct {
		year  int
		month time.Month
		day   int
	}
	tests := []struct {
		name             string
		args             args
		wantStartTimeStr string
		wantEndTimeStr   string
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{year: 2022, month: 6, day: 1},
			wantStartTimeStr: "2022-05-16", wantEndTimeStr: "2022-05-31"},
		{name: "test1", args: args{year: 2022, month: 6, day: 15},
			wantStartTimeStr: "2022-05-16", wantEndTimeStr: "2022-05-31"},
		{name: "test1", args: args{year: 2022, month: 6, day: 16},
			wantStartTimeStr: "2022-06-01", wantEndTimeStr: "2022-06-15"},
		{name: "test2", args: args{year: 2022, month: 1, day: 1},
			wantStartTimeStr: "2021-12-16", wantEndTimeStr: "2021-12-31"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartTimeStr, gotEndTimeStr := GetLastHalfMonth(tt.args.year, tt.args.month, tt.args.day)
			if gotStartTimeStr != tt.wantStartTimeStr {
				t.Errorf("GetLastHalfMonth() gotStartTimeStr = %v, want %v", gotStartTimeStr, tt.wantStartTimeStr)
			}
			if gotEndTimeStr != tt.wantEndTimeStr {
				t.Errorf("GetLastHalfMonth() gotEndTimeStr = %v, want %v", gotEndTimeStr, tt.wantEndTimeStr)
			}
		})
	}
}

func TestGetLastWeek(t *testing.T) {
	start, end := GetLastWeek(-1)
	fmt.Println(start, end)
}
