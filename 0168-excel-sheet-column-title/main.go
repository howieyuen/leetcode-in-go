package _168_excel_sheet_column_title

func convertToTitle(n int) string {
	var s string
	for n > 0 {
		n--
		s = string('A'+n%26) + s
		n /= 26
	}
	return s
}
