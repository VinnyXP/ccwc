package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	//args := os.Args
	//MyFile, err := os.Open(args[1:])
	strC := flag.String("c", "example.txt", "number of bytes") 
	flag.Parse()

	c_str := *strC
	MyFile, err := os.Open(c_str)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer MyFile.Close()

	myFile, _ := MyFile.Stat()

	fmt.Println(myFile.Size(), c_str)
}
