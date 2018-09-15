package rotate

func rotate(s []int, n int) {
	tmp := make([]int, n)
	copy(tmp, s[0:n])
	tmp2 := s[n:len(s)]
	for i, t := range tmp2 {
		s[i] = t
	}
	for i, t := range tmp {
		s[i+len(s)-n] = t
	}
}
