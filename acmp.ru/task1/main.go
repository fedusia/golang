package main

/*
Link: http://acmp.ru/?main=task&id_task=1
Date: 28.05.2015
*/

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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

	r := io.Reader(file)
	lines, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}
	slice := strings.Split(string(lines), " ")
	a, _ := strconv.Atoi(slice[0])
	b, _ := strconv.Atoi(slice[1])
	sum := a + b

	fmt.Println(sum)
}
