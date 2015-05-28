package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
)

var (
	input  = "./INPUT.TXT"
	output = "./OUTPUT.TXT"
)

func main() {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	r,err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(r)) 
}