package helper

import "time"

func WeekdayToInt(day time.Weekday) int {
	return int(day)
}

func ScheduleToDay(day int) string {
	switch day {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	default:
		return "Invalid day"
	}
}
