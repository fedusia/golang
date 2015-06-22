package main

import (
	"log"
	"io/ioutil"
	"strconv"
)

var (
	input = "./INPUT.TXT"
	output = "./OUTPUT.TXT"
)

func main() {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalln(err)
	}
	line := string(bytes)
	a, err := strconv.Atoi(line)
	a = a * a
	result := strconv.Itoa(a)
	ioutil.WriteFile(output, []byte(result), 0644)
}