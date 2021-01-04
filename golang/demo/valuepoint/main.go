package main

import "fmt"

// var a = []int{1, 2, 3}
var a = make([]int, 3, 6)

func change(b []int) {
	fmt.Printf("b 的地址：%p\n", &b)
	fmt.Printf("b 的len：%d\n", len(b))
	fmt.Printf("b 的cap：%d\n", cap(b))
	fmt.Println("b的值为：", b)
	b[0] = 2
	c := append(b, 4)
	fmt.Printf("c 的地址：%p\n", &c)
	fmt.Printf("c 的len： %d\n", len(c))
	fmt.Printf("c 的cap： %d\n", cap(c))
	fmt.Println("c的值为：", c)
	c[0] = 10
	c[1] = 3
	c[2] = 30
	fmt.Printf("修改后的 c 的地址：%p\n", &c)
	fmt.Printf("修改后的 c 的len： %d\n", len(c))
	fmt.Printf("修改后的 c 的cap： %d\n", cap(c))
	fmt.Println("修改后的c的值为：", c)

	fmt.Printf("append之后b 的地址：%p\n", &b)
	fmt.Printf("append之后 b 的len：%d\n", len(b))
	fmt.Printf("append之后 b 的cap：%d\n", cap(b))
	fmt.Println("append之后 a的值为：", a, "append之后b的值为：", b)

}

func main() {
	fmt.Printf("传参之前a 的地址：%p\n", &a)
	fmt.Printf("传参之前 a 的len：%d\n", len(a))
	fmt.Printf("传参之前 a 的cap：%d\n", cap(a))
	fmt.Println("传参之前 a的值为：", a)
	change(a)
	fmt.Printf("传参之后a 的地址：%p\n", &a)
	fmt.Printf("传参之后 a 的len：%d\n", len(a))
	fmt.Printf("传参之后 a 的cap：%d\n", cap(a))
	fmt.Println("传参之后 a的值为：", a)
}
