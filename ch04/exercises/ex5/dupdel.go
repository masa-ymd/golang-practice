package dupdel

func dupdel(s []string) []string {
	var r []string
	var b string
	f := true
	for _, t := range s {
		if f {
			b = t
			f = false
			r = append(r, t)
			continue
		}
		if t == b {
			continue
		}
		r = append(r, t)
		b = t
	}
	return r
}
