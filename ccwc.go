package main

import (
	"fmt"
	"os"
)

func main() {
	MyFile, _ := os.Open("test.txt")
	myFile, _ := MyFile.Stat()

	fmt.Println(myFile.Size())

}