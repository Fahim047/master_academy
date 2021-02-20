package main

import "fmt"

/*
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
*/
/*
func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}
*/
func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Fahad" // defrencing pointer address
	return
}

func main(){
	/*
	sum1 := add(10,30)
	fmt.Println(sum1)
	sum2 := add(5,8)
	fmt.Println(sum2)
	sum3 := add(15,5)
	fmt.Println(sum3)
	sum4 := add(20,30)
	fmt.Println(sum4)
	*/
	/*
	var a, p int
	a, p = rectangle(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
	*/
	age :=21
	name := "Fahim"
	update(&age, &name)
	fmt.Println(age,name)

}