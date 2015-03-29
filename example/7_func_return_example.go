package main

/* Just one more example */

import "fmt"

func add(args ...int) (total int, total2 string) { //Если мы в подписи функции описали, какие переменные она возвращает, то просто используем return
	for i := 0; i < len(args); i++ { //Если в подписи функции мы указали только тип в return пишем что вернуть.
		total += args[i]
	}
	total2 = "Hello world"
	return
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 10))
}
