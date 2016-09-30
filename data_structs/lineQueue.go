package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
)

var (
	queue []string
)

func main() {
	http.HandleFunc("/list", ListQueue)
	http.HandleFunc("/append", AppendQueue)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func InitQueue() {

}

func ListQueue(w http.ResponseWriter, req *http.Request) {
	for _, v := range queue {
		io.WriteString(w, v)
	}
}

func AppendQueue(w http.ResponseWriter, req *http.Request) {
	a := GenData()
	queue = append(queue, a)
	io.WriteString(w, a)

}

/*
func DeleteQueue(q *[]string) string {
	x, a = a[len(a)-1], a[:len(a)-1]
}
*/
func GenData() string {
	rand.Seed(1)
	a := []string{"a", "b", "c"}
	return a[rand.Intn(len(a))]
}
