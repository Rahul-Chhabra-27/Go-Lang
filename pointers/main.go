package main

import "fmt"

func main() {
	fmt.Println("Pointers section")
	number := 20;
	// Ways of creating pointer..
	// 1. var ptr *int = &number;
	// 2. var ptr = &number;
	// 3. ptr := &number;
	ptr := &number;
	fmt.Println("Value of the pointer ",*ptr);
	fmt.Println("Actual Pointer ", ptr);

	*ptr *= 2;
	fmt.Println("Updated Pointer value ", *ptr);
}
