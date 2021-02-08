package main

import "fmt"

func main(){

	x := make(map[string]string)
	
	x["name"]= "Fahimul"
	x["height"]= "5.7"
	x["age"]= "21"
	x["address"]= "Manikganj"

	//all the value
	//fmt.Println(x)
	//specific value
	//fmt.Println(x["name"])

	//delete keys
	delete(x, "height")
	fmt.Println(x)
}