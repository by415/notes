package bytes

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

/// strings 底层使用的bytes里的方法。
func TestBytes(t *testing.T) {
	// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	// The zero value for Buffer is an empty buffer ready to use.
	//type Buffer struct {
	//	buf      []byte // contents are the bytes buf[off : len(buf)]
	//	off      int    // read at &buf[off], write at &buf[len(buf)]
	//	lastRead readOp // last read operation, so that Unread* can work correctly.
	//}

	s1 := "你好hello"
	bt1 := []byte(s1)
	fmt.Println(bt1)
	bf1 := bytes.NewBuffer(bt1)
	fmt.Println("bf", bf1)
	fmt.Println("s1", len(s1))
	fmt.Println("中文支持 runes 长度计算，迭代。都不会出现乱码 ", len(bytes.Runes(bt1)))

	bf2 := bytes.NewBufferString("hello world")
	fmt.Println(bf2)
	fmt.Println("Compare", bytes.Compare(bt1, bf2.Bytes()))

	fmt.Println("Contains", bytes.Contains(bt1, []byte("o")))
	fmt.Println("Equal ", bytes.Equal(bt1, []byte("hello")))
	fmt.Println("Equal ", bytes.Equal(bt1, []byte("你好hello")))
}

// 中文验证 分割
func TestBytes2(t *testing.T) {
	type BTS []byte
	s1 := "你好  hello"
	bt1 := []byte(s1)
	fmt.Println(bytes.ContainsRune(bt1, bytes.Runes([]byte("你ox"))[0]))
	fmt.Println(bytes.ContainsRune(bt1, bytes.Runes([]byte("你ox"))[1]))
	fmt.Println(bytes.ContainsRune(bt1, bytes.Runes([]byte("你ox"))[2]))

	fds := bytes.Fields([]byte("你 sdf sdf23是 个 好人呀"))
	for _, v := range fds {
		fmt.Println(string(v))
	}

	f := func(x rune) bool {
		return string(x) == "你"
	}
	fds2 := bytes.FieldsFunc([]byte("sfs sf你 sdf sdf23你是 个 好人你呀"), f)
	for _, v := range fds2 {
		fmt.Println(string(v))
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(bytes.HasPrefix(bt1, BTS("你")))
	fmt.Println(bytes.HasPrefix(bt1, BTS("x")))

	fmt.Println(string(bytes.TrimPrefix(bt1, BTS("你好"))))

	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(bytes.Count(bt1, BTS("l")))
	rd1 := bytes.NewReader(bt1)
	rd1.ReadRune()

	fmt.Println(strings.EqualFold("GoLang","golang"))
	fmt.Println(strings.EqualFold("golang","GoLang"))
}
