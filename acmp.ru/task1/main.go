package main

/*
Link: http://acmp.ru/?main=task&id_task=1
Date: 18.06.2015
*/

import (
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
	bytes  = make([]byte, 100)
	a      int
	b      int
	sum    int
)

func main() {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	r := io.Reader(file)
	// Читаем все байты из файла
	bytes, err := ioutil.ReadAll(r)
	// Переводим байты в строку
	line := string(bytes)
	//Разбиваем по пробелам строку
	slice := strings.Fields(line)
	//Меняем тип с str на int
	a, err = strconv.Atoi(slice[0])
	b, err = strconv.Atoi(slice[1])
	sum = a + b
	// Меняем тип с int на  str и затем конвертим в байты и записываем в файл
	sum := strconv.Itoa(sum)
	ioutil.WriteFile(output, []byte(sum), 0644)

}
