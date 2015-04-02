package main

import "fmt"

func main() {

	for id := 21; id <= 68; id++ {
		fmt.Printf("Delete from devicemodel where id=%d;\n", id)
	}
}
