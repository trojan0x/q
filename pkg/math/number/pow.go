package number

func Pow(a, r int) int {
	if a == 0 {
		return 0
	}

	if r == 0 {
		return 1
	}

	p := a
	for i := 1; i < r; i++ {
		p = p * a
	}

	return p
}

// ModExp2 returns a^(2^j) mod N
func ModExp2(a, j, N int) int64 {
	if a == 0 {
		return 0
	}

	if j == 0 {
		return int64(a % N)
	}

	p := a
	for i := 0; i < j; i++ {
		p = (p * a) % N
	}

	return int64(p)
}
