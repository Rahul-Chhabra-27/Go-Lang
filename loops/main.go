package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday"}

	fmt.Println(days)

	for i:=0; i < len(days); i++ {
		fmt.Println(days[i]);
	}

	// another way of using loop..
	for i := range days {
		fmt.Println(days[i]);
	}

	// another way of using loop...
	for index,day := range days {
		fmt.Printf("Day no. is %d and day is %s \n", index+1,day);
	}

	// another way of using loop...
	counter := 0;

	for counter <= 10 {
		if counter > 5 {break;}
		fmt.Println(counter);
		counter++;
	}

}
