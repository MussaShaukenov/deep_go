package main

import (
	"fmt"
	"strconv"
	"strings"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	s := fmt.Sprintf("%032b", number) // pad to 32 bits
	bytes := []string{}
	for i := 0; i < 32; i += 8 {
		bytes = append(bytes, s[i:i+8])
	}

	// reverse byte order
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	result := strings.Join(bytes, "")
	endResult, _ := strconv.ParseUint(result, 2, 32)
	return uint32(endResult)
}

func main() {
	a := ToLittleEndian(0x00FF00FF)
	fmt.Println("a=", a)
}
