package io_test

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestIOUlt(t *testing.T) {
	//x, _ := ioutil.ReadDir(".")
	//for _, v := range x {
	//	fmt.Println(v.Name())
	//}
	//
	//ioutil.WriteFile("./tt4",[]byte("xxxxxx"),os.ModePerm)

	//f, _ := os.Open("./test.txt")

	//x, _ := ioutil.ReadAll(f)
	x, _ := ioutil.ReadFile("./test.txt")
	fmt.Println(string(x))

}
