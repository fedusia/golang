package main

/*
Link: http://acmp.ru/?main=task&id_task=2
Date: 21.06.2015
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
	input     = "./INPUT.TXT"
	output    = "./OUTPUT.TXT"
	bytes     = make([]byte, 100)
	a, result int
)

func main() {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	r := io.Reader(file)

	bytes, err = ioutil.ReadAll(r)
	line := string(bytes)
	slice := strings.Fields(line)
	a, err = strconv.Atoi(slice[0])

	for i := 0; i <= a; i++ {
		result += i
	}

	result := strconv.Itoa(result)
	ioutil.WriteFile(output, []byte(result), 0644)

}
