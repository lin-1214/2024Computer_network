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

	f_in, err := os.Open(input_filename)

	stat, err := f_in.Stat()
	check(err)

	fmt.Printf("File size: %v\n", stat.Size())
}