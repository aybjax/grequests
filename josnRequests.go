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
	
	var returnData map[string]interface{}
	resp.JSON(&returnData)
	fmt.Println(returnData)
}