package main

import "fmt"
import "reflect"

func main(){

	// ------1st part------ //
	//directly using array
	var students [3]string

	//students[0]= "Fahim"
	//students[1]= "Rian"
	//students[2]= "Tuli"

	// var_name := array_name[start:quantity]
	//x := students[0:3]
	//fmt.Println(x)
	

	// -------2nd part------ //
	//x := make([]string, 0)
	var fruits []string
	fruits=append(fruits, "Apple", "Banana", "Mango")
	
	fmt.Println(fruits, len(fruits))
	//display the type of variables
	fmt.Printf("%T ----------slice\n", fruits)
	fmt.Printf("%T -------array\n", students)

	//display the type of variables
	//bellow syntax is case-sensitive
	b := reflect.TypeOf(students).Kind().String()
	fmt.Println(b)
	a :=reflect.TypeOf(fruits).Kind().String()
	fmt.Println(a)


}