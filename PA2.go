package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
}

func main() {
	fmt.Printf("Input filename: ")
	var input_filename string = ""
	var output_filename string = ""
	var line int = 1
	fmt.Scanf("%s", &input_filename)

	fmt.Printf("Output filename: ")
	fmt.Scanf("%s", &output_filename)

	f_in, err := os.Open(input_filename)
	f_out, err := os.Create(output_filename)
	check(err)

	scanner := bufio.NewScanner(f_in)
	writer := bufio.NewWriter(f_out)
	for scanner.Scan() {
		var content string = strconv.Itoa(line) + " " + scanner.Text() + "\n"
		writer.WriteString(content)
		writer.Flush()
		line++
	}

	f_in.Close()
	f_out.Close()
}