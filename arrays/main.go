package main

import "fmt"

func main() {

	// In case of int array not initialized value is by default 0...
	// In case of string not initialized value is ''...

	var rollNumbers [4]int
	rollNumbers[0] = 1
	rollNumbers[2] = 2
	rollNumbers[3] = 3

	fmt.Println(rollNumbers)

	var names [3]string

	names[0] = "Rahul"
	names[1] = "Chhabra"

	fmt.Println(names)
	// op: [Rahul Chhabra ] as we can see there is a gap (space) at the end of the array.........
	fmt.Println("Length of the array ", len(names))

	var vegList = [3]string{"Mushroom", "Potato", "Beans"};

	fmt.Println(vegList);
	fmt.Println(len(vegList));
	
}
