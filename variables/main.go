package main

import "fmt"

// statically typed programming language...
// data types

// Can't change the value of const variable....
// Can't use jwtToken := "dcscssfsdf" here..........
const jwtToken string = "fdcdkfndkfd";

func main() {
	// Implicit data type...
	// This will store any value and for future it will behave like the data type you have entered.........
	var username = "chhabrarahul027";

	// can rewrite as
	var username2 string = "Chhabrarahul027";

	var rollnumber int = 10;
	var myGateScore float32 = 720.345543;

	// Doing this will overfolw the data type range....
	// var rollNumber uint8 = 256;
	// fmt.Println(rollNumber);

	fmt.Printf("My username is %s and Roll Number is %d \n", username, rollnumber);
	fmt.Println("My username is chhabrarahul027 and Roll Number is 10");
	fmt.Printf("My username is %s and Roll Number is %d\n", username2,rollnumber);
	fmt.Printf("My gate score is %T type\n", myGateScore);
	
	fmt.Println(myGateScore)

	rollNumber := 123;
	fmt.Println(rollNumber);	
	fmt.Println(jwtToken);

	/// ways of declaring variables in Go
	// 1.var rollNumber int = 123;
   //  2.var rollNumber = 123 Automatically(Implicitly) detect the datatype.
  //   3.rollNumber := 123; // Limitation can't decalre outside the function....

//   Identifier firstLetter Capital -> public
//   Identifier firstLetter Small   -> private 
}	



// var value uint32 = 1000000000
// var value2 int = 10;
// fmt.Println(value)
// var username string = "Rahul Chhabra"
// fmt.Println(username)
// fmt.Printf("Variable is of type: %T \n", username)
// fmt.Printf("Integer printing : %d \n", 100)

