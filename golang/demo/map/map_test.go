package map_test

import (
	"fmt"
	"testing"
)

// TestMap map使用
func TestMap(t *testing.T) {
	mp := make(map[int]int)
	fmt.Println(mp, len(mp))
	mp1 := make(map[int]int, 3)
	fmt.Println(mp1, len(mp1))

	mp2 := map[int]int{}
	fmt.Println(mp2, len(mp2))
	mp3 := map[int]int{1: 3, 2: 4, 3: 5}
	fmt.Println(mp3, len(mp3))

	mp[4] = 3
	fmt.Println(mp, len(mp))

	for k, v := range mp3 {
		fmt.Println(k, v)
	}
	delete(mp3, 2)
	fmt.Println()
	for k, v := range mp3 {
		fmt.Println(k, v)
	}
	_, ok := mp3[55]
	fmt.Println(ok)
}

type User struct {
	name string
	age  int
}

// TestMapStruct 验证map struct
func TestMapStruct(t *testing.T) {
	ur := User{
		name: "n",
		age:  1,
	}

	mp := map[int]User{}
	// 不能通过 map 引用直接修改
	//mp[1].age = 3

	mp[1] = ur
}
