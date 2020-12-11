package _649_dota2_senate

func predictPartyVictory(senate string) string {
	var r, d []int
	for i := range senate {
		if senate[i] == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	n := len(senate)
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+n)
		} else {
			d = append(d, d[0]+n)
		}
		r = r[1:]
		d = d[1:]
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}
