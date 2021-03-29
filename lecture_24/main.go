package main

import (
	"fmt"
	"os"
)

func main() {

	// show the directory of current file
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(dir)
	// -------------------------------------------------------------------------------------
	
	// creating file 
	// posf, err := os.Create("fahim.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer posf.Close()

	// n, err := posf.Write([]byte("This is a test file. I am Fahad."))
	// fmt.Println(n, err)
	// --------------------------------------------------------------------------------------
	
	// calling createFile function
	isErr := CreateFile("rian.txt", "This is second test file. Rian is my bajan.")
	fmt.Println(isErr)
}

// creating file using function
func CreateFile(filename, content string) bool {

	posf, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	// closing file 
	// defer will compile at the end 
	defer posf.Close()

	_, err = posf.Write([]byte(content))
	if err != nil {
		fmt.Println(err.Error())
	}

	return true
}
