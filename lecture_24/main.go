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
	// isErr := CreateFile("rian.txt", "This is second test file. Rian is my bajan.")
	// fmt.Println(isErr)

	// fInfo, err := os.Stat("master_academy.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(fInfo.IsDir())
	// fmt.Println(fInfo.ModTime().Date())
	// fmt.Println(fInfo.ModTime().Clock())
	// fmt.Println(fInfo.Name())
	// fmt.Println(fInfo.Size())
	// -------------------------------------------------------------------------------------
	// How to make a folder
	// err := os.Mkdir("Test_Folder", 764)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// -------------------------------------------------------------------------------------
	// base := filepath.Base(dir)
	// relativePath := filepath.Join("Test_Folder")
	// absolutePath, err := filepath.Abs("Test_Folder")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// newPath := filepath.Join(absolutePath, "..", "..", "lecture_24_test")

	// fmt.Println(base)
	// fmt.Println(relativePath)
	// fmt.Println(absolutePath)
	// fmt.Println(newPath)

	os.Mkdir(`F:\GO_WORKSPACE\src\Master_Academy\lecture_24_test2`, 777)
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