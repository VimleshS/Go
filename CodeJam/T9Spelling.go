package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var rawKeysChars = map[int]string{}
var t9HashMap = map[string][]int{}

func main() {
	rawKeysChars[1] = ""
	rawKeysChars[2] = "ABC"
	rawKeysChars[3] = "DEF"
	rawKeysChars[4] = "GHI"
	rawKeysChars[5] = "JKL"
	rawKeysChars[6] = "MNO"
	rawKeysChars[7] = "PQRS"
	rawKeysChars[8] = "TUV"
	rawKeysChars[9] = "WXYZ"
	rawKeysChars[0] = " "
	t9HashMap = invertMap(rawKeysChars)
	ReadFile("T9.txt")
}

/*
  Returns slice mapped to chars
  Character = [Digit, RepeatCount]
  A= [2,1] B=[2,2] H=[4,2]
*/
func invertMap(rawMap map[int]string) map[string][]int {
	charMappedToIntSlice := make(map[string][]int)
	for key, value := range rawMap {
		//charMappedToIntSlice[value] = append(charMappedToIntSlice[value], key)
		chars := []rune(value)
		for pos, char := range chars {
			charMappedToIntSlice[string(char)] = []int{key, pos + 1}
		}
	}
	return charMappedToIntSlice
}

func ReadFile(filename string) {
	file, _ := os.Open(filename)
	filereader := bufio.NewReader(file)
	var idx int = 0
	firstTime := true
	for {
		line, err := filereader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file: ", err)
			}
			break
		}
		line = line[:len(line)-2]
		if firstTime == true {
			//no of cases here is redundant... (first line)
			firstTime = !firstTime
			continue
		}
		idx += 1
		//Get the keyStrokes and Print
		getT9data(idx, line)
	}
}

func getT9data(idx int, data string) {
	frstTime := true
	prevInt := -1

	var t9string string
	for _, charinrune := range data {
		char := strings.ToUpper(string(charinrune))

		t := t9HashMap[char]
		s := strconv.Itoa(t[0])

		if frstTime == false {
			if prevInt == t[0] {
				t9string = t9string + " "
			}
			prevInt = t[0]
		}
		buff := strings.Repeat(s, t[1])
		t9string = t9string + buff

		if frstTime {
			prevInt = t[0]
			frstTime = !frstTime
		}
	}
	fmt.Println("Case #", idx, " ", t9string)
}
