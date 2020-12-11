package _483_smallest_good_base

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)
	for m := int(math.Log2(float64(num))); m >= 1; m-- {
		l, r := 2, int(math.Pow(float64(num), 1/float64(m)))+1
		for l < r {
			k := (l + r) / 2
			sum := 1
			for j := 0; j < m; j++ {
				sum = sum*k + 1
			}
			if sum == num {
				return strconv.Itoa(k)
			} else if sum < num {
				l = k + 1
			} else {
				r = k
			}
		}
	}
	return strconv.Itoa(num - 1)
}

func format(a []int) string {
	var res string
	for _, c := range a {
		res += strconv.Itoa(c)
	}
	return res
}

/*
    123
  x 456
--------
    738
   615
+ 492
--------
  56088
*/
func multiply(a, b []int) []int {
	var res []int
	for i := len(a) - 1; i >= 0; i-- {
		var carry int
		var tmp []int
		for j := len(b) - 1; j >= 0; j-- {
			r := a[i]*b[j] + carry
			carry = r / 10
			r %= 10
			tmp = append(tmp, r)
		}
		if carry > 0 {
			tmp = append(tmp, carry)
		}
		reverse(tmp)
		tmp = appendZero(tmp, len(a)-i-1)
		res = add(res, tmp)
	}
	return res
}

/*
2|56088
  -----
8|16
  -----
  0|0
    ---
   4|8
     ---
     4|8
       -
*/
func divide(a []int) []int {
	var res []int
	var carry int
	for _, c := range a {
		c += carry * 10
		if c < 2 {
			carry = 1
			continue
		}
		carry = c % 2
		res = append(res, c/2)
	}
	return res
}

func less(a, b []int) int {
	if len(a) < len(b) {
		return 1
	} else if len(a) > len(b) {
		return -1
	} else {
		for i := range a {
			if a[i] < b[i] {
				return 1
			} else if a[i] > b[i] {
				return -1
			} else {
				continue
			}
		}
		return 0
	}
}

func add(a, b []int) []int {
	var res []int
	i, j := len(a)-1, len(b)-1
	var carry int = 0
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		r := a[i] + b[j] + carry
		carry = r / 10
		r %= 10
		res = append(res, r)
	}
	for ; i >= 0; i-- {
		r := a[i] + carry
		carry = r / 10
		r %= 10
		res = append(res, r)
	}
	for ; j >= 0; j-- {
		r := b[j] + carry
		carry = r / 10
		r %= 10
		res = append(res, r)
	}
	if carry > 0 {
		res = append(res, carry)
	}
	reverse(res)
	return res
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func appendZero(a []int, c int) []int {
	for c > 0 {
		a = append(a, 0)
		c--
	}
	return a
}
