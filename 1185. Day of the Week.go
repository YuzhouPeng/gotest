package leetcode_go

func dayOfTheWeek(day int, month int, year int) string {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if month < 3 {
		month += 12
		year -= 1
	}

	var c int = year / 100
	year = year % 100
	var w int = (c/4 - 2*c + year + year/4 + 13*(month+1)/5 + day - 1) % 7
	return days[(w+7)%7]
}
