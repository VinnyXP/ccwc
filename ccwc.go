package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
	"bufio"
)

//Implementation of the c flag in wc
func c_flag(textFile string) string {
	if textFile == "" {
		return "Error: No file path provided"
	}

	MyFile, err := os.Open(textFile)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer MyFile.Close()

	myFile, _ := MyFile.Stat()
	return strconv.FormatInt(myFile.Size(), 10) + " " + textFile
}

//Implementation of the l flag in wc
func l_flag(textFile string) string {
	if textFile == "" {
		return "Error: No file path provided"
	}

	MyFile, err := os.Open(textFile)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer MyFile.Close()

	scanner := bufio.NewScanner(MyFile)
	scanner.Split(bufio.ScanLines)

	var count int
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return ""
	}
	return strconv.Itoa(count) + " " + textFile
}

func w_flag(textFile string) string{
	if textFile == "" {
		return "Error: No file path provided"
	}

	MyFile, err := os.Open(textFile)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return ""
	}

	defer MyFile.Close()

	scanner := bufio.NewScanner(MyFile)
	scanner.Split(bufio.ScanWords)

	var count int
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return ""
	}
	return strconv.Itoa(count) + " " + textFile
}

func m_flag(textFile string) string{
	if textFile == "" {
		return "Error: No file path provided"
	}

	MyFile, err := os.Open(textFile)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return ""
	}

	defer MyFile.Close()

	scanner := bufio.NewScanner(MyFile)
	scanner.Split(bufio.ScanRunes)

	var count int
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return ""
	}
	return strconv.Itoa(count) + " " + textFile
}

func main() {
	//cli options
	strC := flag.String("c", "", "number of bytes") 
	strL := flag.String("l", "", "number of lines")
	strW := flag.String("w", "", "number of words")
	strM := flag.String("m", "", "number of character")
	flag.Parse()

	if *strC != "" && *strL != "" && *strW != ""{
		//If the user uses more than one flag
		fmt.Println("Error: Please provide only one of the flags: -c, - l, -w, or -m.")
	} else if *strC != "" {
		fmt.Println(c_flag(*strC))
	} else if *strL != "" {
		fmt.Println(l_flag(*strL))
	} else if *strW != "" {
		fmt.Println(w_flag(*strW))
	} else if *strM != "" {
		fmt.Println(m_flag(*strM))
	} else {
		//If no flags were given
		fmt.Println("Error: No valid flag value provided")
	}
}
