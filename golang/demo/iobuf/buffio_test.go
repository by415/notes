package iobuf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestIOBufRead(t *testing.T) {
	f, _ := os.Open("input.txt")
	defer f.Close()
	//bf := bufio.NewReader(f)
	//bs := make([]byte, 4)
	//
	//for true {
	//	n, _ := bf.Read(bs)
	//	if n <= 0 {
	//		break
	//	}
	//	fmt.Println(string(bs))
	//}

	//bf := bytes.NewBuffer(make([]byte, 1024))
	//r := bufio.NewReaderSize(f, 20)
	//r.Read(bf.Bytes())
	//fmt.Println(bf.String())

	bf := bufio.NewReader(f)

	for true {
		n, err := bf.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(string(n))
	}
}

func TestIOBufWrite(t *testing.T) {
	f, _ := os.Open("input.txt")
	defer f.Close()
	//
	//bf := bytes.NewBuffer(make([]byte, 1024))
	//r := bufio.NewReaderSize(f, 20)
	//r.Read(bf.Bytes())
	//fmt.Println(bf.String())

	fo, _ := os.Create("out3.txt")
	defer fo.Close()

	w := bufio.NewWriter(fo)
	w.ReadFrom(f)
	w.Flush()
}

func TestIOBufReadWrite(t *testing.T) {
	f, _ := os.Open("input.txt")
	defer f.Close()

	fo, _ := os.Create("out4.txt")
	defer fo.Close()

	w := bufio.NewReadWriter(bufio.NewReader(f), bufio.NewWriter(fo))
	s, _ := w.ReadString('\n')
	fmt.Println(s)
	w.Write([]byte(s))
	w.Flush()
	//w.Flush()
}
