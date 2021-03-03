package byteorder_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestGetOrder(t *testing.T) {
	var a = 0x11223344
	// fmt.Printf("%x\n", *(*byte)(unsafe.Pointer(&a)))
	if *(*int8)(unsafe.Pointer(&a)) == 0x44 {
		fmt.Println("small")
	} else {
		fmt.Println("big")
	}
	// return *(*byte)(unsafe.Pointer(&a)) == 0x44
}
