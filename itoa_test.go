package itoa_test

import (
	"strconv"
	"testing"

	"github.com/kayabe/itoa"
)

var ranges = [...]uint32{
	0, 9,
	10, 99,
	100, 999,
	1000, 9999,
	10000, 99999,
	100000, 999999,
	1000000, 9999999,
	10000000, 99999999,
	100000000, 200000000, // 999999999
	1000000000, 1200000000, // 4294967295
}

var numbers = [...]struct {
	A   uint32
	Len uint
}{
	{1234567890, 10},
	{1254557590, 10},
	{123456789, 9},
	{12345678, 8},
	{1234567, 7},
	{123456, 6},
	{12345, 5},
	{1234, 4},
	{123, 3},
	{12, 2},
	{1, 1},
}

const numbersLen = len(numbers)

func TestItoaU32(t *testing.T) {
	t.Parallel()

	var buf = make([]byte, 10)
	var x uint = 1

	for i := 0; i < len(ranges); i += 2 {
		start := ranges[i]
		end := ranges[i+1]

		name := "1e" + strconv.FormatUint(uint64(x), 10)

		t.Run(name, func(pt *testing.T) {
			for j := start; j <= end; j++ {
				itoa.U32toa(j, x-1, buf)

				c := strconv.FormatUint(uint64(j), 10)

				if string(buf[:x]) != c {
					pt.Logf("FormatUint[%d] %s", i, c)
					pt.Logf("U32toa    [%d] %s", i, buf[:x])
					pt.Fail()
					break
				}
			}
		})

		x++
	}
}

func TestItoaUnsafeU32(t *testing.T) {
	t.Parallel()

	var buf = make([]byte, 10)
	var x uint = 1

	for i := 0; i < len(ranges); i += 2 {
		start := ranges[i]
		end := ranges[i+1]

		name := "1e" + strconv.FormatUint(uint64(x), 10)

		t.Run(name, func(pt *testing.T) {
			for j := start; j <= end; j++ {
				itoa.U32toaUnsafe(j, x-1, buf)

				c := strconv.FormatUint(uint64(j), 10)

				if string(buf[:x]) != c {
					pt.Logf("FormatUint[%d] %s", i, c)
					pt.Logf("U32toa    [%d] %s", i, buf[:x])
					pt.Fail()
					break
				}
			}
		})

		x++
	}
}

func BenchmarkU32toa(b *testing.B) {
	buf := make([]byte, 10)
	for i := 0; i < b.N; i++ {
		itoa.U32toa(numbers[i%numbersLen].A, numbers[i%numbersLen].Len-1, buf)
	}
}

func BenchmarkU32toaUnsafe(b *testing.B) {
	buf := make([]byte, 10)
	for i := 0; i < b.N; i++ {
		itoa.U32toaUnsafe(numbers[i%numbersLen].A, numbers[i%numbersLen].Len-1, buf)
	}
}

func BenchmarkFormatUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatUint(uint64(numbers[i%numbersLen].A), 10)
	}
}
