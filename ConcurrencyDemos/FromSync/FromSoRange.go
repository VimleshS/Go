package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	intInputChan := make(chan int, 50)
	wg := new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(intInputChan, wg)
	}
	for i := 1; i < 51; i++ {
		intInputChan <- i
	}
	close(intInputChan)
	wg.Wait()
	fmt.Println("Existing Main App... ")
	//panic("---------------")
}

func worker(input chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range input {
		fmt.Printf("range %d\n", v)
		time.Sleep(100 * time.Millisecond)
	}
}
