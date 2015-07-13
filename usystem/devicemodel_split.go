package main

import "fmt"
import "os"
import "log"
import "encoding/csv"

func main() {
	path_to_file := "/home/fedusia/devicemodel.csv"
	file, err := os.Open(path_to_file)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	id := 21
	for _, value := range lines {
		fmt.Printf("INSERT INTO \"devicemodel\" VALUES(%[1]d,'%[2]s', %[3]s, '%[4]s', %[5]s, '%[6]s', '%[7]s', '%[8]s');\n", id, value[0], value[1], value[2], value[3], value[4], value[5], value[6])
		id++
	}
}
