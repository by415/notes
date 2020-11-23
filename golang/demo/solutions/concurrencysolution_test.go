package solutions

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

/// channel 关闭，需要
func TestChClose(t *testing.T) {
	ch := make(chan int)
	close(ch)
	fmt.Println(<-ch)

	if _, ok := <-ch; ok {
		fmt.Println(ok)
	} else {
		fmt.Println("closed")
	}
}

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

// TestSelect testSelect 随机选择执行
func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	func1 := func(ch1 chan int) {
		for true {
			select {
			case ch1 <- 0:
			case ch1 <- 1:
			}
		}
	}
	go func1(ch1)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch1)
	}
}

func TestEndNotice(t *testing.T) {
	GenerateIntA := func(done chan struct{}) chan int {
		ch := make(chan int)
		go func() {
		Label:
			for true {
				select {
				case ch <- rand.Int():

				case <-done: /// 增加一路监听，监听推出通知信号
					break Label
				}
			}
		}()
		return ch
	}

	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("NumGoroutine = ", runtime.NumGoroutine())
}

/// 生成器简单版本
func TestSimpleGenerate(t *testing.T) {
	GenerateIntA := func() chan int {
		ch := make(chan int, 10)
		go func() {
			for true {
				ch <- rand.Int()
			}
		}()
		return ch
	}

	ch := GenerateIntA()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/// 增强版本生成器,使用推出通知信号
func TestStandGenerate(t *testing.T) {
	GenerateIntA := func(done chan struct{}) chan int {
		ch := make(chan int, 10)
		go func() {
			fmt.Println("GenerateIntA")
		Label:
			for true {
				select {
				case <-done:
					break Label
				case ch <- rand.Int():
				}
			}
			close(ch)
		}()
		return ch
	}
	GenerateIntB := func(done chan struct{}) chan int {
		ch := make(chan int, 10)
		go func() {
			fmt.Println("GenerateIntB")
		Label:
			for true {
				select {
				case <-done:
					break Label
				case ch <- rand.Int():
				}
			}
			close(ch)
		}()
		return ch
	}
	//
	done := make(chan struct{})
	Gen := func(done chan struct{}) chan int {
		ch := make(chan int, 20)
		send := make(chan struct{})
		go func() {
		Label:
			for true {
				/// 使用select的扇入技术，处理生产者生产慢，消费这消费块的问题
				select {
				case <-done:
					send <- struct{}{}
					send <- struct{}{}
					break Label
				case ch <- <-GenerateIntA(send):
				case ch <- <-GenerateIntB(send):
				}
			}
			close(ch)
		}()
		return ch
	}

	fmt.Println(runtime.NumGoroutine())
	ch := Gen(done)

	// timer start
	go func() {
		time.Sleep(time.Second * 1)
		done <- struct{}{}
		fmt.Println("xxxxxxxx")
	}()
	for {
		v, ok := <-ch
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	// timer end

	/// 计数器完成关闭 start
	//for i := 0; i < 50; i++ {
	//	fmt.Println(<-ch)
	//}
	//close(done)
	/// 计数器完成关闭 end
}

// 管道实现
func TestPipe(t *testing.T) {
	pipe := func(in chan int) chan int {
		out := make(chan int)
		go func() {
			for v := range in {
				out <- v + 1
			}
			close(out)
		}()
		return out
	}

	in := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	out := pipe(pipe(pipe(in)))
	for v := range out {
		fmt.Println(v)
	}
}

/// 管道的关闭时机？

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 计算0-100和
// task 封装实现
// 工作任务
type task struct {
	begin  int
	end    int
	result chan<- int
}

// 任务执行，计算beg-end和
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum = sum + i
	}
	t.result <- sum
}
func InitTask(taskChan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10

	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		task := task{b, e, r}
		taskChan <- task
	}

	if mod != 0 {
		task := task{high + 1, p, r}
		taskChan <- task
	}
	close(taskChan)
}

func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

func DistributeTask(taskChan <-chan task, wait *sync.WaitGroup, resultChan chan int) {
	for v := range taskChan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(resultChan)
}
func ProcessResult(resultChan chan int) int {
	sum := 0
	for v := range resultChan {
		sum = sum + v
		fmt.Println(sum)
	}
	return sum
}

func TestTask(t *testing.T) {
	// 任务通道
	taskChan := make(chan task, 10)
	// 结果通道
	resultChan := make(chan int, 10)

	wait := &sync.WaitGroup{}

	/// 创建task并写入task通道
	go InitTask(taskChan, resultChan, 100)

	go DistributeTask(taskChan, wait, resultChan)

	sum := ProcessResult(resultChan)
	fmt.Println("sum = ", sum)
}

func DistributeTask2(taskChan <-chan task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask2(taskChan, done)
	}
}

func ProcessTask2(taskChan <-chan task, done chan struct{}) {
	for t := range taskChan {
		t.do()
	}
	done <- struct{}{}
}

func CloseResult(done chan struct{}, resultChan chan int, works int) {
	for i := 0; i < works; i++ {
		<-done
	}
	close(done)
	close(resultChan)
}

// 固定work 的工作池
const NUMBER = 10

// 与上边的差别。不是结束通知。更加
func TestTaskNumber(t *testing.T) {
	workers := NUMBER

	taskChan := make(chan task, 10)

	chanResult := make(chan int, 10)

	done := make(chan struct{}, 10)

	go InitTask(taskChan, chanResult, 100)

	DistributeTask2(taskChan, workers, done)

	go CloseResult(done, chanResult, workers)

	sum := ProcessResult(chanResult)

	fmt.Println(sum)
}
