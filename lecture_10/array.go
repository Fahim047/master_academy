package main

import "fmt"

func main(){

	//var studentName [3]string
	//studentName[0]= "Fahim"
	//studentName[1]= "Rian"
	//studentName[2]= "Tuli"
	//fmt.Println(studentName)
	//fmt.Println(len(studentName))
	//fmt.Println(studentName[2])
	
	studentName :=[...]string{"Fahim", "Rian", "Tuli", "Rabbe"}
	fmt.Println(studentName[1])
	fmt.Println(len(studentName))
}