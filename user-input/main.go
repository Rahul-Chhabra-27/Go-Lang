package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome user"
	fmt.Println(welcome)

	rating := bufio.NewReader(os.Stdin);
	fmt.Println("Enter the Rating for the pizza ");

	// Comma ok || err ok operator...
	input, _ := rating.ReadString('\n');

	fmt.Println("Thanks for rating ", input);
	// Here the type of the input is string.........
	// let's see how to convert it into number....
	fmt.Printf("The type of the input is %T", input);
	// fmt.Println("Thanks for rating", input,input,input,input);
}
