package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now();
	time.Sleep(time.Second * 4);
	fmt.Println("Start")
	//time.Sleep(time.Second * 4);
	fmt.Println("End")
	//time.Sleep(time.Second * 4);

	fmt.Print(time.Since(now));
}
