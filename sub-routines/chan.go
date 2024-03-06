package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()

	smokeSignal := make(chan bool)

	defer func() {
		fmt.Println(time.Since(currentTime))
	}()

	evilNinja := "Raha"
	go attack(evilNinja, smokeSignal)
	go yess(smokeSignal)
	fmt.Println(<- smokeSignal); // recieving the message..........
	fmt.Println(<- smokeSignal); // recieving the message..........
	//time.Sleep(time.Second * 2);
	//fmt.Println("jj")
}
func yess(smokeSignal chan bool) {
	time.Sleep(time.Second * 2);
	smokeSignal <- false;
}
func attack(target string, attacked chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println("Throwing ninjs starts at ", target)
	attacked <- true // sending the message.......
}
