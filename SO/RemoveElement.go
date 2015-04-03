package main

import (
	//	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func main() {

	urlList := []string{"test", "abc", "def", "ghi"}
	remove := []string{"abc", "test"}

	for idx := len(urlList) - 1; idx >= 0; idx-- {
		url := urlList[idx]
		for _, val := range remove {
			if url == val {
				urlList = append(urlList[:idx], urlList[idx+1:]...)
				continue
			}
		}
	}
	fmt.Println(urlList)
	b := make([]byte, 2048) // adjust buffer size to be larger than expected stack
	n := runtime.Stack(b, false)
	s := string(b[:n])
	fmt.Println(s)
	//	Find(urlList, "abc")
}

func Find(a []string, str string) {
	s := strings.Join(a, "")
	fmt.Println(strings.Index(s, str))
	//	bytes.Index(s, []byte(str))

	//	fmt.Println(strings.Join(a, ""))
	//	//	bt := make([]byte, 50)
	//	//	fmt.Println(search[0])
	//	//	btnew := copy(bt, search[0])
	//	//	fmt.Println(btnew)

	//	n := len(a) - 1
	//	for i := 0; i < len(a); i++ {
	//		n += len(a[i])
	//	}

	//	b := make([]byte, n)
	//	bp := copy(b, a[0])
	//	for _, s := range a[1:] {
	//		//		bp += copy(b[bp:], "")
	//		bp += copy(b[bp:], s)
	//	}
	//	//	fmt.Printf("%T  %v\n", b, b)
	//	//	fmt.Printf("%s\n", string(b))

}
