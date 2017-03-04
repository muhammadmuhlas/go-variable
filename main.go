package main

import "fmt"

func main() {

	myInteger := 10
	myString := "golang"
	myDouble := 123.54
	myBoolean := true

	//The %T represent of the Type and the %v represent the default value of its variable
	//Only Work with "Printf"
	fmt.Printf("%T %v \n", myInteger, myInteger)
	fmt.Printf("%v \n", myString)
	fmt.Printf("%v \n", myDouble)
	fmt.Printf("%v \n", myBoolean)
}



