package main

import (
	//	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var (
	queue []string
	x     string
	q     *[]string
)

func main() {
	q = InitQueue()
	//	http.HandleFunc("/delete", DeleteQueue)
	http.HandleFunc("/list", ListQueue)
	http.HandleFunc("/append", AppendQueue)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func InitQueue() *[]string {
	return &queue
}

func list(q *[]string) string {
	return strings.Join(*q, ",")
}

func ListQueue(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, list(q))
}

func AppendQueue(w http.ResponseWriter, req *http.Request) {
	a := GenData()
	*q = append(*q, a)
	io.WriteString(w, a)
}

func GenData() string {
	rand.Seed(1)
	a := []string{"a", "b", "c"}
	return a[rand.Intn(len(a))]
}

/*
func DeleteQueue(w http.ResponseWriter, req *http.Request) {
	a := Remove(q)
	io.WriteString(w, a)
}

func Remove(q *[]string) string {
	x, q = q[len(q)-1], q[:len(q)-1]
	return x
}
*/
