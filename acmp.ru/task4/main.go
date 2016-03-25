package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strconv"
)

const (
	b = 9
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
	c := b - a
	d := fmt.Sprintf("%d%d%d",a, b, c)
	err = ioutil.WriteFile(output, []byte(d), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}