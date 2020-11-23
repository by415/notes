package solutions

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/////////////////////////////////////////////////////////////////
// 1.使用channel 处理同步问题
// TestRunTine2
func TestRunTine2(t *testing.T) {
	d := make(chan string)
	go func() {
		fmt.Println("func")
		time.Sleep(time.Second * 3)
		d <- "done"

	}()
	fmt.Println(<-d)
	fmt.Println("end in main func")
}

/////////////////////////////////////////////////////////////////
// 2.使用sync.WaitGroup 处理同步问题
// TestRunTine1
func TestRunTine1(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("func")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end in main func")
}

/////////////////////////////////////////////////////////////////
// 2.使用带缓存的 chan 处理同步问题  基于 chal 特性
// 从无数据channel读数据会阻塞
// 往无空间channel写数据会阻塞
// TestRunTine3
func TestRunTine3(t *testing.T) {
	chs := make(chan int, 10)
	for i := 1; i < cap(chs); i++ {
		time.Sleep(time.Second * 1)
		go func(i int) {
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			chs <- i
		}(i)
	}
	for i := 1; i < cap(chs); i++ {
		<-chs
	}
	fmt.Println("end Run")
}

/// 生产者消费者
// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(time.Second * 2)
		out <- 1 * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestProCons(t *testing.T) {
	chPro := make(chan int, 12)
	go Producer(3, chPro)
	go Producer(5, chPro)
	go Consumer(chPro)

	time.Sleep(time.Second * 10)
}

type Student struct {
	age int
	ID  int
}

type Class struct {
	ID    int
	Count int
}

/// M 生产者 N消费者
func Pro(ID, inner chan<- Class) {

}
func TestMN(t *testing.T) {
	//ch1 := make(chan int)
	//select {}
}
