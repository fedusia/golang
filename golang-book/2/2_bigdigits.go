package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/* Объявляем срез срезов */
var bigdigits = [][]string{
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "2222 "},
	{"  333 ", " 3   3", "     3", "   33 ", "     3", " 3   3", "  333 "}, /*  Все элементы среза должны иметь   */
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ", "   4  "}, /*  Все элементы среза должны иметь   */
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},        /*  должны иметь одинаковую  		  */
	{" 666  ", "6     ", "6     ", "66666 ", "6   6 ", "6   6 ", " 666  "}, /*  длинну строки. 					  */
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},        /*  В противном случае вывод поплывет */
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}

func main() {
	/*if len(os.Args) > 0 {
		fmt.Printf("try --help or -h\n")
		os.Exit(1)
	}
*/
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Printf("usage: %s [-b|--bar] <whole number>\n", filepath.Base(os.Args[0]))
		fmt.Printf("-b --bar draw an underbar and an overbar\n")
		os.Exit(1)
	}

	if len(os.Args) == 2 && (os.Args[1] == "-b" || os.Args[1] == "--bar") {
		stringOfDigits := os.Args[1]
		for row := range bigdigits[0] {
			line := ""
			for column := range stringOfDigits {
				digit := stringOfDigits[column] - '0' /* Не понятная строка еще раз перечитать главу и разобраться как сравниваются строки в Go */
				if 0 <= digit && digit <= 9 {
					line += bigdigits[digit][row] + "  "
				} else {
					log.Fatal("Invalid whole number")
				}
			}
			fmt.Println(line)
		}
	}
}