package _731_my_calendar_ii

import (
	`fmt`
	`testing`
)

func TestMyCalendarTwo_Book(t *testing.T) {
	c := Constructor()
	days := [][2]int{
		{36, 41},
		{28, 34},
		{40, 46},
		{10, 18},
		{4, 11},
		{25, 34},
		{36, 44},
		{32, 40},
		{34, 39},
		{40, 49},
	}
	for _, day := range days {
		fmt.Println(c.Book(day[0], day[1]))
	}
}

/*
{10, 20}, // true
{50, 60}, // true
{10, 40}, // true
{5, 15},  // false
{5, 10},  // true
{25, 55}, // true
*/
