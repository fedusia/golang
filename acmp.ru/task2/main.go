package main

/*
Link: http://acmp.ru/?main=task&id_task=2
Date: 21.06.2015
*/

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	input     = "./INPUT.TXT"
	output    = "./OUTPUT.TXT"
	result int
)

func main() {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalln(err)
	}
	line := string(bytes)
	slice := strings.Fields(line)
	a, err := strconv.Atoi(slice[0])
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i <= a; i++ {
		result += i
	}

	result := strconv.Itoa(result)
	ioutil.WriteFile(output, []byte(result), 0644)

}
