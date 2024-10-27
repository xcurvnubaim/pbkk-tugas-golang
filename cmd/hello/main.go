package main

import (
	"os"
	"strconv"
)

func main() {
	if(len(os.Args) < 2) {
		println("Usage : main <number>")
		return
	}

	i, err := strconv.Atoi(os.Args[1]); 
	
	if err != nil {
		panic("input must be a number")
	} 

	if i == 42 {
		println("Hello, Universe!")
	} else {
		println(os.Args[1])
	}
}