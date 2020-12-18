package _389_find_the_difference

func findTheDifference(s string, t string) byte {
	var z byte
	for i := range t {
		z += t[i]
	}
	for i := range s {
		z -= s[i]
	}
	return z
}
