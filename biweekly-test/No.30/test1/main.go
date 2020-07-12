package test1

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {
	var months = map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "07",
		"Jul": "06",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	splits := strings.Split(date, " ")
	year := splits[2]
	month := months[splits[1]]
	var day = make([]byte, 0, 2)
	for i := range splits[0] {
		if splits[0][i] <= '9' && splits[0][i] >= '0' {
			day = append(day, splits[0][i])
		} else {
			break
		}
	}
	if len(day) == 1 {
		day = append(day, day[0])
		day[0] = '0'
	}
	return fmt.Sprintf("%s-%s-%s", year, month, string(day))
}
