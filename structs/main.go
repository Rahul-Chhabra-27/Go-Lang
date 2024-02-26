package main

import "fmt"

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func main() {
	fmt.Println("Structs");

	rahul := User{"Rahul","xyz@123",true,1};
	fmt.Println(rahul)
	fmt.Printf("User details are %+v: \n", rahul);

	fmt.Printf("Name is %v and Email is %v \n", rahul.Name, rahul.Email);
}