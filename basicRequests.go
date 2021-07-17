package main

import (
	"fmt"
	"log"

	"github.com/levigross/grequests"
)

func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	
	if err != nil {
		log.Fatalln("unable to make request:", err)
	}

	// fmt.Printf("%#v\n", resp)
	// fmt.Println(error(nil))
	fmt.Println(resp.String())
}