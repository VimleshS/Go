package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	msg string
}

func main() {
	//	c := boring("Vimlesh")

	timeout := time.After(5 * time.Second)
	//	for {
	//		select {
	//		case s := <-c:
	//			fmt.Println(s.msg)
	//		case <-timeout:
	//			fmt.Println("Existing Switch ....")
	//			return
	//		}
	//	}

	quit, bye := make(chan bool), make(chan bool)
	fanIn(boring("Vimlesh"), quit, bye, timeout)
	<-quit
	fmt.Println("Code Clean up...")
	bye <- true
	fmt.Println("Existing Main Function")
}

func fanIn(c <-chan Message, quit, bye chan bool, timeout <-chan time.Time) {
	go func() {
		for {
			select {
			case msg := <-c:
				fmt.Println(msg.msg)
			case <-timeout:
				fmt.Println("TimedOut.....")
				quit <- true
			case <-bye:
				fmt.Println("Bye.....")
			}
		}
	}()
}

func boring(msg string) <-chan Message {
	c := make(chan Message)

	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i)}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func routine(str string, c chan int) {
	for {
		t := <-c
		fmt.Printf("%s %d", str, t)
		fmt.Println()
		time.Sleep(100 * time.Millisecond)
		c <- t + 1
	}
}
