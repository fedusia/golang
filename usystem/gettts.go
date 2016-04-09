package main

import "net/http"
import "fmt"


func main() {
	resp, err  := http.Get("http://tts.naukanet.ru")
	req, err := http.NewRequest("GET", "http://tts.naukanet.ru", nil)
	req.SetBasicAuth("", "")
	resp, err := resp.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

//func NewRequest(method, urlStr string, body io.Reader) (*Request, error)