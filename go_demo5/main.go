package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func testA() {
	for i := 0; i < 8; i++ {
		fmt.Println("testA" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	wg.Done()
}

func testB() {
	for i := 0; i < 10; i++ {
		fmt.Println("testB" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test_1() {
	wg.Add(1)
	go testA()
	go testB()
	for i := 0; i < 5; i++ {
		fmt.Println("main" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	wg.Wait()
}

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello !", i)
}

func test_2() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

func main() {
	// test_1()
	fmt.Println(runtime.NumCPU())
	test_2()
}
