package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hello")
	client := new(http.Client)
	r, e := client.Get("https://google.ca")
	defer client.CloseIdleConnections()

	if e != nil {
		fmt.Println(e.Error())

	}
	defer r.Body.Close()

	s, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(s))

}
