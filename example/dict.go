package main

import "fmt"
import "os"


func main() {

	dir, err := os.Open("./") // For read access.
	if err != nil {
		fmt.Println("qeqw", err)
	}
	fi, _ := dir.Readdir(-1)	
	for _ , fi := range fi {
		fmt.Println(fi.Name())
	}
}



