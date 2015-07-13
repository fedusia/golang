package main

/*
Link: http://acmp.ru/?main=task&id_task=1
Date: 18.06.2015
*/

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	input  = "./INPUT.TXT"
	output = "./OUTPUT.TXT"
)

func main() {
	// Читаем все байты из файла
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalln(err)
	}
	// Переводим байты в строку
	line := string(bytes)
	//Разбиваем по пробелам строку
	slice := strings.Fields(line)
	//Меняем тип с str на int
	a, err := strconv.Atoi(slice[0])
	if err != nil {
		log.Fatalln(err)
	}
	b, err := strconv.Atoi(slice[1])
	if err != nil {
		log.Fatalln(err)
	}
	sum := a + b
	// Меняем тип с int на  str и затем конвертим в байты и записываем в файл
	result := strconv.Itoa(sum)
	ioutil.WriteFile(output, []byte(result), 0644)

}
