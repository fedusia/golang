package main

import "net/http"
import "log"
import "fmt"
import "html"


func main() {
	http.HandleFunc("/fedusia", handler)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err) }
}

func handler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}