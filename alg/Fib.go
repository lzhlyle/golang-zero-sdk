package alg

func FibArr(n int) (res []int) {
	if n < 0 {
		return
	}

	a, b := 0, 1
	res = append(res, a, b)
	for i := 0; i < n; i++ {
		a, b = b, a+b
		res = append(res, b)
	}
	return
}
