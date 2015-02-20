package main

import (
	"fmt"
	// _ "image/gif"
	// _ "image/png"
	//	"bufio"
	"image"
	//	"image/jpeg"
	_ "image/png"
	//	"io/ioutil"
	"log"
	"os"
)

func main() {

	//	dat, err := ioutil.ReadFile("download.png")
	//	fmt.Println(string(dat))

	reader, err := os.Open("download.png")
	if err != nil {
		log.Fatal(err)

	}
	defer reader.Close()

	//	breader := bufio.NewReader(reader)

	//	l, _ := reader.Stat()
	//	x := l.Size()
	//	b := make([]byte, x)
	//	_, err2 := breader.Read(b)
	//	if err2 != nil {
	//		fmt.Println(err2)
	//	}
	//	fmt.Printf("Size of byte %X  \n  %v \n", b, b)

	//	i, _ := png.Decode(reader)
	//	w, _ := os.Create("down.jpg")

	//	o := &jpeg.Options{}
	//	o.Quality = 100

	//	jpeg.Encode(w, i, o)
	//	fmt.Println(i)

	img, _, err1 := image.Decode(reader)
	if err1 != nil {
		log.Fatal(err1)
	}
	bounds := img.Bounds()

	//	fmt.Println(str)
	//	fmt.Println("----------------------------")
	//fmt.Printf("%T \n", img)

	//	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			normR := r / 256
			normG := g / 256
			normB := b / 256
			fmt.Printf("%X  %6X %6X \n", normR, normG, normB)
		}
	}

	// Print the results.
	//	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	//	for i, x := range histogram {
	//		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	//	}

}
