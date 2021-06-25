package togroup

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	t := float64(x) / 2

	for {
		v := (t + float64(x)/t) / 2

		if t-v < 0.001 {
			return int(v)
		}

		t = v
	}
}
