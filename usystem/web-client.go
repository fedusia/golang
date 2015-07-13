package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.Client{}

	request, err := http.NewRequest("GET", "http://noc.naukanet.ru/sa/managedobject/6021", nil)
	request.SetBasicAuth("login", "pass")
	resp, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
