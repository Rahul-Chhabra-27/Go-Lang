package main

import (
	"fmt"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {

	response, err := http.Get(url)
	// Response is just the pointer...... *http.Response.
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // caller's responsibility is to close the connection.... 
	fmt.Printf("Type of response is %T", response);
}
