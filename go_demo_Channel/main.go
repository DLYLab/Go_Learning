package main

import (
	"fmt"
	"sync"
	// "time"
)

var wg sync.WaitGroup

func fn1(ch chan int){
	for i:=0; i<1000; i++ {
		ch <- i + 1
		fmt.Print("in", i+1, " ")
		// time.Sleep(time.Millisecond * 100)
	}
	close(ch)
	wg.Done()
}

func fn2(ch chan int) {
	for v := range ch {
				fmt.Print("out", v, " ")
		// time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func main(){
	allch := make(chan int, 100)
	wg.Add(1)
	go fn1(allch)
	wg.Add(1)
	go fn2(allch)
	wg.Wait()
	fmt.Println("ok")
}
