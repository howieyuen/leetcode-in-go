package _017_Letter_Combinations_of_a_Phone_Number

func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) > 0 {
		combineNext("", digits, &ans)
	}
	return ans
}

var digit2letter = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func combineNext(cur string, digits string, ans *[]string) {
	if len(digits) == 0 {
		*ans = append(*ans, cur)
	} else {
		letters := digit2letter[digits[0:1]]
		for i := 0; i < len(letters); i++ {
			combineNext(cur+letters[i:i+1], digits[1:], ans)
		}
	}
}

func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var digit2Letter = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var ans []string
	var dfs func(index int, cur []byte)
	dfs = func(index int, cur []byte) {
		if index == len(digits) {
			ans = append(ans, string(cur))
			return
		}
		sub := digit2Letter[digits[index]]
		for j := range sub {
			cur = append(cur, sub[j])
			dfs(index+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0, nil)
	return ans
}
