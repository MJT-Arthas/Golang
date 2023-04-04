package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

//
//import (
//	"fmt"
//	"sync"
//)
//
//var wg sync.WaitGroup
//
//func hello(i int) {
//	defer wg.Done() // goroutine结束就登记-1
//	fmt.Println("hello", i)
//}
//func main() {
//	for i := 0; i < 10; i++ {
//		wg.Add(1) // 启动一个goroutine就登记+1
//		go hello(i)
//	}
//	wg.Wait() // 等待所有登记的goroutine都结束
//
//	for i := 0; i < 5; i++ {
//		go func() {
//			fmt.Println(i)
//		}()
//	}
//}

// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
//func Producer() chan int {
//	ch := make(chan int, 2)
//	// 创建一个新的goroutine执行发送数据的任务
//	go func() {
//		for i := 0; i < 10; i++ {
//			if i%2 == 1 {
//				ch <- i
//			}
//		}
//		close(ch) // 任务完成后关闭通道
//	}()
//
//	return ch
//}
//
//// Consumer 从通道中接收数据进行计算
//func Consumer(ch chan int) int {
//	sum := 0
//	for v := range ch {
//		sum += v
//	}
//	return sum
//}
//
//func main() {
//	ch := Producer()
//
//	res := Consumer(ch)
//	fmt.Println(res) // 25
//
//}
//
//var (
//	x int64
//
//	wg sync.WaitGroup // 等待组
//	m  sync.Mutex     // 互斥锁
//)
//
//// add 对全局变量x执行5000次加1操作
//func add() {
//	for i := 0; i < 5000; i++ {
//		m.Lock()
//		x = x + 1
//		m.Unlock()
//	}
//	wg.Done()
//}
//
//func main() {
//	wg.Add(2)
//
//	go add()
//	go add()
//
//	wg.Wait()
//	fmt.Println(x)
//}

var wg sync.WaitGroup

func Producer() chan int {
	defer wg.Done()
	wg.Add(1)
	ch := make(chan int, 100)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 100; i++ {

			ch <- rand.Intn(math.MaxInt64)

		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

func Consumer(ch chan int, resultChan chan int) {
	defer wg.Done()
	wg.Add(1)
	for v := range ch {
		resultChan <- cal(v)
	}

}

func cal(x int) int {
	var res int = 0
	for x != 0 {
		res += x % 10
		x /= 10
	}
	return res
}
func main() {
	resultChan := make(chan int, 10)
	ch := Producer()
	for i := 0; i < 24; i++ {
		go Consumer(ch, resultChan)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for res := range resultChan {
		fmt.Println(res)
	}
}
