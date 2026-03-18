package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fun_1() {
	start := time.Now().UnixMilli()
	for num := 1; num <= 240000; num++ {
		flag := true //假设是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(num)
		}
	}
	end := time.Now().UnixMilli()
	fmt.Println("普通的方法耗时=", end-start)
}

func fun_2(n int) {
	for num := (n-1)*60000 + 1; num <= n*60000; num++ {
		flag := true //假设是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(num)
		}
	}
	wg.Done()
}

func test1() {
	start := time.Now().UnixMilli()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go fun_2(i)
	}
	wg.Wait()
	end := time.Now().UnixMilli()
	fmt.Println("协程的方法耗时=", end-start)
}

func main() {
	fun_1()
	test1()
}
