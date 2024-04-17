package itoa

// U32toaUnsafe ...
//	n is supposed to be number of digits - 1
func U32toaUnsafe(u uint32, n uint, b []byte) {
	switch n {
	case 0:
		setOneUnsafe(0, b, u)
	case 1:
		setTwoUnsafe(0, b, u)
	case 2:
		setTwoUnsafe(0, b, u/10)
		setOneUnsafe(2, b, u%10)
	case 3:
		setTwoUnsafe(0, b, u/100)
		setTwoUnsafe(2, b, u%100)
	case 4:
		setTwoUnsafe(0, b, u/1_000)
		setTwoUnsafe(2, b, u%1_000/10)
		setOneUnsafe(4, b, u%10)
	case 5:
		setTwoUnsafe(0, b, u/10_000)
		setTwoUnsafe(2, b, u%10_000/100)
		setTwoUnsafe(4, b, u%100)
	case 6:
		setTwoUnsafe(0, b, u/100_000)
		setTwoUnsafe(2, b, u%100_000/1_000)
		setTwoUnsafe(4, b, u%1_000/10)
		setOneUnsafe(6, b, u%10)
	case 7:
		setTwoUnsafe(0, b, u/1_000_000)
		setTwoUnsafe(2, b, u%1_000_000/10_000)
		setTwoUnsafe(4, b, u%10_000/100)
		setTwoUnsafe(6, b, u%100)
	case 8:
		setTwoUnsafe(0, b, u/10_000_000)
		setTwoUnsafe(2, b, u%10_000_000/100_000)
		setTwoUnsafe(4, b, u%100_000/1_000)
		setTwoUnsafe(6, b, u%1_000/10)
		setOneUnsafe(8, b, u%10)
	case 9:
		setTwoUnsafe(0, b, u/100_000_000)
		setTwoUnsafe(2, b, u%100_000_000/1_000_000)
		setTwoUnsafe(4, b, u%1_000_000/10_000)
		setTwoUnsafe(6, b, u%10_000/100)
		setTwoUnsafe(8, b, u%100)
	}
}

// U64toaUnsafe ...
//	partial work
func U64toaUnsafe(u uint64, n uint, b []byte) {
	if n>>32 == 0 {
		U32toaUnsafe(uint32(u), n, b)
		return
	}
}
