package main

import "os"

func main() {
	if(len(os.Args) < 2) {
		println("Usage : main <number>")
		return
	}	

	if os.Args[1] == "42" {
		println("Hello, Universe!")
	} else {
		println(os.Args[1])
	}
}