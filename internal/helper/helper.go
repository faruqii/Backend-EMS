package helper

import (
	"strconv"
	"time"
)

func WeekdayToInt(day time.Weekday) int {
	return int(day)
}

func WeekdayToStr(day time.Weekday) string {
	return strconv.Itoa(int(day))
}


