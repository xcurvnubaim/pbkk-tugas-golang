package main

import "os"

func main() {
	if(len(os.Args) < 2) {
		println("Usage : reverse <string1> <string2> <string3> ...")
		return
	} else if (len(os.Args) < 4) {
		println("Input at least 3 strings")
		return
	}

	
	for i, str := range os.Args {
		if i == 0 {
			continue
		}
		var reversed string
		for j := len(str) - 1; j >= 0; j-- {
			reversed += string(str[j])
		}
		print(reversed + " ")	
	}
}