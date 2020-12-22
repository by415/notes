package syncmap

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var syncmap = new(sync.Map)
var mp = make(map[int]int)

/// 线程安全的读写。需要同步操作保证有数据
func TestSyncMap(t *testing.T) {
	cd := sync.NewCond(&sync.Mutex{})
	for i := 1; i < 20; i++ {
		go func(i int) {
			syncmap.Store(i, i)
			//fmt.Println("store", i, i)
		}(i)
		time.Sleep(1*time.Second)
		go func(i int) {
			cd.L.Lock()
			defer cd.L.Unlock()
			for v, ok := syncmap.Load(i); ok == false || v == 0 || v == nil; {
				//fmt.Println("wait", i, ok, v)
				cd.Wait()
			}
			v, _ := syncmap.Load(i)
			fmt.Println(v)
		}(i)
	}
	time.Sleep(2 * time.Second)
}

// 条件变量的锁
func TestSyncMap2(t *testing.T) {
	cd := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 40; i++ {
		go func(i int) {
			cd.L.Lock()
			mp[i] = i
			cd.L.Unlock()
		}(i)
		go func(i int) {
			cd.L.Lock()
			fmt.Println(mp[i])
			cd.L.Unlock()
		}(i)
	}
}

// 条件变量的锁
func TestSyncMap3(t *testing.T) {
	cd := sync.NewCond(&sync.Mutex{})
	for i := 1; i < 20; i++ {
		go func(i int) {
			cd.L.Lock()
			for mp[i] != 0 {
				cd.Wait()
			}
			mp[i] = i
			cd.Broadcast()
			cd.L.Unlock()
		}(i)
		go func(i int) {
			cd.L.Lock()
			for mp[i] == 0 {
				cd.Wait()
			}
			fmt.Println(mp[i])
			cd.Broadcast()
			cd.L.Unlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
}

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知，阻塞当前 goroutine

			// do something. 这里仅打印
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 进入 Wait 阻塞状态
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发下一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 1)
	cond.Broadcast() // 1 秒后下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 执行完毕
}
