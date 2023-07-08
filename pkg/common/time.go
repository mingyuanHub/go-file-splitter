package common

import (
	"strings"
	"time"
)

func GetNowSecond() int64 {
	return time.Now().Unix()
}

func GetNowMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func DataToTimeSecond(day string) int64 {
	day = strings.Replace(day, ".", "-", -1)

	list := strings.Split(day, "-")
	day = ""
	for _, item := range list {
		if len(item) == 1 {
			item = "0" + item
		}
		day += item + "-"
	}
	day = strings.Trim(day, "-")

	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("2006-01-02", day, loc)
	return  tt.Unix()
}

func DataToTimeMillSecond(day string) int64 {
	return DataToTimeSecond(day) * 1000
}
