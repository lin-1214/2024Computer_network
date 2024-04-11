package main

import(
	"fmt"
	"os"
)

func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
}

func main() {
	fmt.Printf("Input filename: ")
	input_filename := ""
	fmt.Scanf("%s", &input_filename)

	contentBytes, err := os.ReadFile(input_filename)
	check(err)

	content := string(contentBytes)
	fmt.Printf("File content:\n%s\n", content)
}