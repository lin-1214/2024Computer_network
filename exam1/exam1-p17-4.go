package main

import(
	"fmt"
	"os"
	"strings"
	"bufio"
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
	f_out, err := os.Create("whatever.txt")
	check(err)

	contentBytes, err := os.ReadFile(input_filename)
	check(err)

	content := strings.ToUpper(string(contentBytes)) + "\n"
	writer := bufio.NewWriter(f_out)
	writer.WriteString(content)
	writer.Flush()
}