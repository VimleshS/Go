// +build Omit

package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("Vimlesh.Sharma@synerzip.com" + ":" + "newPass1")))
	//fmt.Println(base64.StdEncoding.DecodeString([]byte("21KQV3Z4NL77XVISQNWTFADN1" + ":" + "80Twrr3oDxCJ90SvYrMzy5kNBhvGM0S4O/tyWZ+/TT8")))
	//	auth := "21KQV3Z4NL77XVISQNWTFADN1" + ":" + "80Twrr3oDxCJ90SvYrMzy5kNBhvGM0S4O//tyWZ+//TT8"

	auth := "51FHLL8SATEVYG283J7IJHI0L" + ":" + "JzxVfiLsgemVnYsbAiAQfNOegZ+B2/vE1PkWRzGXNYI"

	fmt.Println(base64.StdEncoding.EncodeToString([]byte(auth)))

	TestArraySlice()
	//panic("ooooooo")
	//CheckSliceModification()
	//CheckMap()
	//CheckSliceIteration()
	//if k, e := CheckPanicAndRecoverUsage(10, 10); e != nil {
	//	fmt.Println(e)
	//} else {
	//	fmt.Println(k)
	//}

}

// end Omit

func CheckPanicAndRecoverUsage(i, j int) (k int, err error) {
	defer func() {
		if e := recover(); e != nil {
			//err = fmt.Errorf("%v", e)  or
			err = errors.New("MyError")
		}
	}()
	//divide by zero causes panic which is handeled
	//by defered anonymous function using recover
	k = i / j
	return k, nil
}

func CheckSliceIteration() {
	//s := "XabYcZ"
	s := "XαβYγZ"
	for i, c := range s {
		fmt.Println(i, c, string(c))
	}
	fmt.Println("---------------------------------------------")
	for index := range s { // String per character iteration .
		char, size := utf8.DecodeRuneInString(s[index:])
		fmt.Println(index, char, string(char), size)
	}
}

func AddToSlice(slice []string) []string {
	return append(slice, "X")
}

func TestArraySlice() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	fmt.Printf("%-4d %-4d %v\n", len(s), cap(s), s)
	//s = s[3:]
	//fmt.Printf("%-4d %-4d %v\n", len(s), cap(s), s)
	s = append(s, "A")
	fmt.Printf("%-4d %-4d %v\n", len(s), cap(s), s)
	//s = AddToSlice(s)
	//fmt.Printf("%-4d %-4d %v\n", len(s), cap(s), s)
}

type Product struct {
	name  string
	price float64
}

func (product Product) String() string {
	return fmt.Sprintf("%s (%.2f)", product.name, product.price)
}

type Point struct {
	x, y, z int
}

//func (point Point) String() string {
//	return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
//}

func CheckMap() {
	nameForPoint := make(map[Point]string)
	p := Point{1, 1, 1}
	nameForPoint[p] = "x"
	nameForPoint[Point{54, 158, 89}] = "y"
	fmt.Println(nameForPoint)
	x := nameForPoint[p]
	fmt.Println(x)

}

func CheckSliceModification() {
	//original value will be modified
	//..way1
	//products := []*Product{&Product{"A", 1}, &Product{"B", 2}}
	//..way2
	//products := []*Product{{"X", 2}, {"Y", 5}}

	//no modification to original will be made
	products := []Product{{"X", 2}, {"Y", 5}}
	fmt.Println(products)
	for _, product := range products {
		product.price += 0.50
	}
	fmt.Println(products)
}
