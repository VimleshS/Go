// +build OMIT

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mylock sync.Mutex

func main() {
	go DoPrint("Hello")
	go DoPrint("World")

	fmt.Scanln()
}

// endmain OMIT

func DoPrint(msg string) {
	for i := 0; i < 5; i++ {
		mylock.Lock()
		fmt.Println(fmt.Sprintf(" %s %s", "", msg))
		//time.Sleep(100 * time.Millisecond)
		mylock.Unlock()
		runtime.Gosched()
	}
}
