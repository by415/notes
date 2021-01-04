package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func Sleep(i int) {
	cmd := exec.Command("sleep", "30")
	cmd.Start()
	cmd.Wait()
	fmt.Println("end sleep:", i)
}

func Sleep2(i int) {
	time.Sleep(30 * time.Second)
	fmt.Println("end sleep:", i)
}

func main() {
	fmt.Println("pid: ", os.Getpid())

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Println("socket error:", err)
		return
	}

	SockaddrInet4 
	syscall.Connect(fd, )

	err = syscall.Listen(fd, 5)
	if err != nil {
		fmt.Println(err)
		return
	}

	nfd, sa, err := syscall.Accept(fd)
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 100
	for i := 0; i < count; i++ {
		go Sleep2(i)
	}

	time.Sleep(50 * time.Second)

}
