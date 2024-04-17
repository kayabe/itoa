package itoa

// U32toa ...
//	n is supposed to be number of digits - 1
func U32toa(u uint32, n uint, b []byte) {
	switch n {
	case 0:
		setOne(0, b, u)
	case 1:
		setTwo(0, b, u)
	case 2:
		setTwo(0, b, u/10)
		setOne(2, b, u%10)
	case 3:
		setTwo(0, b, u/100)
		setTwo(2, b, u%100)
	case 4:
		setTwo(0, b, u/1_000)
		setTwo(2, b, u%1_000/10)
		setOne(4, b, u%10)
	case 5:
		setTwo(0, b, u/10_000)
		setTwo(2, b, u%10_000/100)
		setTwo(4, b, u%100)
	case 6:
		setTwo(0, b, u/100_000)
		setTwo(2, b, u%100_000/1_000)
		setTwo(4, b, u%1_000/10)
		setOne(6, b, u%10)
	case 7:
		setTwo(0, b, u/1_000_000)
		setTwo(2, b, u%1_000_000/10_000)
		setTwo(4, b, u%10_000/100)
		setTwo(6, b, u%100)
	case 8:
		setTwo(0, b, u/10_000_000)
		setTwo(2, b, u%10_000_000/100_000)
		setTwo(4, b, u%100_000/1_000)
		setTwo(6, b, u%1_000/10)
		setOne(8, b, u%10)
	case 9:
		setTwo(0, b, u/100_000_000)
		setTwo(2, b, u%100_000_000/1_000_000)
		setTwo(4, b, u%1_000_000/10_000)
		setTwo(6, b, u%10_000/100)
		setTwo(8, b, u%100)
	}
}

// U64toa ...
//	partial work
func U64toa(u uint64, n uint, b []byte) {
	if n>>32 == 0 {
		U32toa(uint32(u), n, b)
		return
	}
}
