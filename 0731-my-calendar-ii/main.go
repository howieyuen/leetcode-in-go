package _731_my_calendar_ii

import (
	`fmt`
)

type MyCalendarTwo struct {
	once  [][2]int
	twice [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		once:  [][2]int{},
		twice: [][2]int{},
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, days := range this.twice {
		if days[0] < end && start < days[1] { // 与二次预定交叉
			fmt.Println()
			fmt.Println(this.once)
			fmt.Println(this.twice)
			return false
		}
	}
	for _, days := range this.once {
		if days[0] < end && start < days[1] { // 首次出现预定交叉
			newStart := max(days[0], start)
			newEnd := min(days[1], end)
			this.twice = append(this.twice, [2]int{newStart, newEnd}) // 加入交集
		}
	}
	this.once = append(this.once, [2]int{start, end})
	fmt.Println()
	fmt.Println(this.once)
	fmt.Println(this.twice)
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*type Calendar struct {
	days [][2]int
}

var overlaps *Calendar

func NewCalendar() *Calendar {
	overlaps = &Calendar{days: [][2]int{}}
	return overlaps
}

func (c *Calendar) book(s, e int) bool {
	for _, days := range c.days {
		if max(days[0], s) < min(days[1], e) {
			return false
		}
	}
	c.days = append(c.days, [2]int{s, e})
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type MyCalendarTwo struct {
	days [][2]int
}

func Constructor() MyCalendarTwo {
	NewCalendar()
	return MyCalendarTwo{
		days: [][2]int{},
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, days := range overlaps.days {
		if days[0] < end && start < days[1] {
			return false
		}
	}
	for _, days := range this.days {
		if days[0] < end && start < days[1] {
			newStart := max(days[0], start)
			newEnd := min(days[1], end)
			overlaps.days = append(overlaps.days, [2]int{newStart, newEnd})
		}
	}
	this.days = append(this.days, [2]int{start, end})
	return true
}*/

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
