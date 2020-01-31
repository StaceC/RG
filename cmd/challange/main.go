package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/StaceC/RG/pkg/code"
)

func main() {

	if len(os.Args) <= 0 {
		fmt.Println("Please specify a file name")
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		encoded, err := code.Encode(scanner.Text(), true)
		if err != nil {
			panic(err)
		}
		fmt.Println(encoded)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
