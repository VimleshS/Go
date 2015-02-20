package main

import (
	"fmt"
	"github.com/koyachi/go-nude"
	"log"
)

func main() {
	imagePath := "download.jpg"

	isNude, err := nude.IsNude(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("isNude = %v\n", isNude)
}
