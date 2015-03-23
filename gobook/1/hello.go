package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "World!"
	fmt.Println("Arguments:", os.Args) /* Print all arguments to stdout */
	if len(os.Args) > 1 {              /* os.Args[0] - имя команды «hello» или «hello.go» */
		fmt.Println("Argument count:", len(os.Args)) /* Print arguments count to stdout */
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}