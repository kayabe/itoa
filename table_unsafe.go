package itoa

import "unsafe"

func setOneUnsafe(index int, buf []byte, value uint32) {
	*(*byte)(unsafe.Pointer(&buf[index])) = byte(value) + '0'
}

func setTwoUnsafe(index int, buf []byte, value uint32) {
	*(*[2]byte)(unsafe.Pointer(&buf[index])) = table[value]
}
