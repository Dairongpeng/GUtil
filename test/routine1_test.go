package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/**
 * 信号量控制，分别打印大写和小写字符三次
 */
func TestTree_Insert(t *testing.T) {
	// 分配一个逻辑处理器给调度器使用、给这个函数传入 1，是通知调度器只
	// 能为该程序使用一个逻辑处理器。
	runtime.GOMAXPROCS(1)

	// wg 用来等待程序完成
	// 计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
				fmt.Println()
			}
		}
	}()

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
				fmt.Println()
			}
		}
	}()

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
