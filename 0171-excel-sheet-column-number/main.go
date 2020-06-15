package _171_excel_sheet_column_number

// 26进制
func titleToNumber(s string) int {
	var num = 0
	for i := range s {
		num = num*26 + int(s[i]-'A'+1)
	}
	return num
}

