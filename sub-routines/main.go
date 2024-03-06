package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	defer func() {
		fmt.Println(time.Since(currentTime))
	}()
	evilNinjas := []string{"Tommy", "Bobby", "Atas", "Raha"}
	for _, ninja := range evilNinjas {
		go attack(ninja)
	}
}

func attack(target string) {
	fmt.Println("Throwing ninjs starts at ", target)
	time.Sleep(time.Second)
}
