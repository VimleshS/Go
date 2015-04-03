//package main

//import "fmt"
// +build omit

//func main() {
//	jobs := make(chan int, 5)
//	done := make(chan bool)
//	go func() {
//		for {
//			j, more := <-jobs
//			if more {
//				fmt.Println("received job", j)
//			} else {
//				fmt.Println("received all jobs")
//				done <- true
//				return
//			}
//		}
//	}()
//	for j := 1; j <= 3; j++ {
//		jobs <- j
//		fmt.Println("sent job", j)
//	}
//	close(jobs)
//	fmt.Println("sent all jobs")
//	<-done
//}

package main

import (
	"fmt"
	//"runtime"
	//"sync"
	"time"
)

func main() {
	intInputChan := make(chan int, 50)
	done := make(chan bool)

	for i := 0; i < 2; i++ {

		go worker(intInputChan, done)
	}
	for i := 1; i < 51; i++ {
		fmt.Printf("Inputs. %d \n", i)
		intInputChan <- i
	}
	close(intInputChan)
	<-done

	fmt.Println("Existing Main App... ")
	panic("---------------")
}

//endmain omit

func worker(input chan int, done chan bool) {
	//defer func() {
	//	fmt.Println("Executing defer..")
	//	wg.Done()
	//}()

	for {
		select {
		case intVal, ok := <-input:
			if ok {
				fmt.Printf("%d  %v\n", intVal, ok)
				time.Sleep(100 * time.Millisecond)
			} else {
				input = nil
				done <- true
				close(done)
				return
			}
			//default:
			//	runtime.Gosched()
		}
	}

}
