package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main()  {
	resp, _ := http.Get("https://centraal.academy")
	page, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", page)
}