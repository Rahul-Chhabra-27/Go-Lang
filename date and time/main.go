package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to date and time")

	presentTime := time.Now();

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2001 13:12:12 Monday"));
}
