package main

import (
	"fmt"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	return (number>>24)&0xFF |
		((number >> 8) & 0xFF00) |
		((number << 8) & 0xFF0000) |
		((number << 24) & 0xFF000000)
}

func main() {
	a := ToLittleEndian(0x00FF00FF)
	fmt.Println("a=", a)
}
