package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Thread struct {
	result int;
	message string;
} 
// ways of creating functions......
func add(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func Greeter() {
	fmt.Println("Greeter");
}
func proAdder(values ...int) Thread {
	result := 0;

	for _, value := range values {
		result += value;
	}
	return Thread{result, "Your functions invoked successfully......"};
}

func adder(array[3] int) int {
	result := 0;
	for _ , val := range(array) {
		result += val;
	}
	return result;
}
func main() {
	func () {
		fmt.Println("IIFE 1 Immediately Invoked........");
	}()
	Greeter();
	io := bufio.NewReader(os.Stdin)

	a, _ := io.ReadString('\n')
	b, _ := io.ReadString('\n')

	valueOne, _ := strconv.ParseInt(strings.TrimSpace(a), 10, 32)
	valueTwo, _ := strconv.ParseInt(strings.TrimSpace(b), 10, 32)

	fmt.Printf("Result is : %d\n", add(int(valueOne), int(valueTwo)))

	fmt.Printf("Thread struct is : %+v\n",  proAdder(1,2,3,4,5,6,7,8,10));
	
	fmt.Println("Result of Adder is : ", adder([3]int{1,2,3}))
	func () {
		fmt.Println("IIFE 2 Immediately Invoked.........");
	}()
}
