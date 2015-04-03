// +build OMIT

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	intInputChan := make(chan int, 50)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(intInputChan, &wg)
	}
	for i := 1; i < 51; i++ {
		//fmt.Printf("Inputs. %d \n", i)
		intInputChan <- i
	}
	close(intInputChan)
	wg.Wait()
	fmt.Println("Existing Main App... ")
	//panic("---------------")
}

// endmain OMIT

func worker(input chan int, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Executing defer..")
		wg.Done()
	}()

	for {
		select {
		case intVal, ok := <-input:

			if !ok {
				input = nil
				return
			}
			fmt.Printf("%d  %v\n", intVal, ok)
			time.Sleep(100 * time.Millisecond)

		default:
			runtime.Gosched()
		}
	}

}

// endmain OMIT
