package study_case

// 读写互斥锁
/*
有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，
这种场景下使用读写锁是更好的一种选择
读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来



读写锁分为两种：读锁和写锁。
当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，
如果是获取写锁就会等待；
当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
*/

import (
	"fmt"
	"sync"
	"time"
)

var (
	x3  int64
	wg3 sync.WaitGroup
	// lock3 sync.Mutex
	rwlock sync.RWMutex
)

// write 获取写锁，对全局变量x3进行写操作
// 该函数会锁定资源，执行写操作后释放锁
func write() {
	// 获取写锁，阻止其他goroutine读写操作
	rwlock.Lock()
	x3 = x3 + 1
	time.Sleep(10 * time.Millisecond)
	// 释放写锁，允许其他goroutine访问
	rwlock.Unlock()
	wg3.Done()
}

// read 获取读锁，对全局变量x3进行读操作
// 该函数会锁定资源进行读取，读取完成后释放锁
func read() {
	// 获取读锁，允许多个goroutine同时读取
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	// 释放读锁
	rwlock.RUnlock()
	wg3.Done()
}

// WRLockCase 读写锁使用示例函数
// 启动100个写goroutine和1000个读goroutine来演示读写锁的性能优势
// 通过统计执行时间来展示读写锁在读多写少场景下的效率
func WRLockCase() {
	start := time.Now()

	// 启动100个写操作的goroutine
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go write()
	}

	// 启动1000个读操作的goroutine
	for i := 0; i < 1000; i++ {
		wg3.Add(1)
		go read()
	}

	// 等待所有goroutine执行完成
	wg3.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start)) // 打印耗时，与普通互斥锁相比，读写锁在这个场景下会快一些
}
