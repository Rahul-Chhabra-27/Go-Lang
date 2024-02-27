package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	io:= bufio.NewReader(os.Stdin);

	loginCount,_ := io.ReadString('\n'); 

	input,_ := strconv.ParseInt(strings.TrimSpace(loginCount),10,64);

	var result string

	if input > 10 {
		loggedIn := true;
		fmt.Println(loggedIn);
		result = "Rahul loggedin"
	} else if input < 10 {
		result = "not logged in"
	} else {
		result = "equal";
	}
	fmt.Println(result);

	// Weird Part.........
	if num := 12; num > 11 {
		fmt.Println(num);
	}
}
