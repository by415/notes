package io_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

/// 从文件读取
func TestIORead(t *testing.T) {
	f, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal("err")
	}
	defer f.Close()
	info, _ := f.Stat()
	fmt.Println(f.Name(), info.Name(), info.IsDir(), info.Mode(), info.ModTime(), info.Size(), info.Sys())

	// 读取中文乱码，byte没有识别，被裁断
	//for {
	//	dt := make([]byte, 4)
	//	count, _ := f.Read(dt)
	//	if count == 0 || err == io.EOF {
	//		break
	//	}
	//	fmt.Println(count, string(dt))
	//}

	bf := bytes.NewBuffer(make([]byte, 0))
	for {
		dt := make([]byte, 4)
		count, _ := f.Read(dt)
		if count == 0 || err == io.EOF {
			break
		}
		//fmt.Println(count, string(dt))
		bf.Write(dt[:count])
	}
	fmt.Println(bf.String())

}

func TestIOWrite(t *testing.T) {
	f, err := os.Create("./out.txt")
	if err != nil {
		log.Fatal("err")
	}
	defer f.Close()

	n, err := f.WriteString("this is a hahah ")
	if err != nil {
		log.Fatal("err")
	}
	fmt.Println(n)
}

// 新建文件夹
func TestMkDir(t *testing.T) {
	os.Mkdir("tt",os.ModeDir)
	os.Mkdir("tt/sfaf",os.ModeDir)
}

