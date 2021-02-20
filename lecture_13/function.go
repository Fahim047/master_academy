package main

import "fmt"

//***method-1***
func add(x int, y int) int{
	z := x+y
	return z
}
//***method-2***
func add2(x , y int) int{
	z := x+y
	return z
}
//***method-3***
func add3(x int, y int) (z int){
	z = x+y
	return z
}
//***method-4***
func add4(x int, y int) (z int){
	z = x+y
	return
}

func main(){
	sum1 := add(10,30)
	fmt.Println(sum1)
	sum2 := add(5,8)
	fmt.Println(sum2)
	sum3 := add(15,5)
	fmt.Println(sum3)
	sum4 := add(20,30)
	fmt.Println(sum4)
}