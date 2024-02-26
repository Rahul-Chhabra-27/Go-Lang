package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our Pizza House..")

	reader := bufio.NewReader(os.Stdin);

	fmt.Println("Please provide the rating : ");
	input, _ := reader.ReadString('\n');

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64);
	fmt.Println("Thanks for the Rating ", input);

	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println("Added 1 to your rating : ", numRating + 1);
	}
}
