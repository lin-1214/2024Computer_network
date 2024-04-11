package main

import (
	"fmt"
	"os"
	"bufio"
	"net"
)

func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
}

func main() {
	
	// build connection
	conn, errc := net.Dial("tcp", "140.112.42.221:12000")
	check(errc)
	defer conn.Close()
	
	// get input
	fmt.Printf("Input filename: ")
	input_filename := ""
	fmt.Scanf("%s", &input_filename)

	f_in, err := os.Open(input_filename)
	contentBytes, err := os.ReadFile(input_filename)
	check(err)

	stat, err := f_in.Stat()
	check(err)

	fmt.Printf("Send the file size first: %v\n", stat.Size())
	content := string(contentBytes)
	
	writer := bufio.NewWriter(conn)
	writer.WriteString(fmt.Sprintf("%d\n", stat.Size()))
	writer.Flush()

	writer.WriteString(content)
	writer.Flush()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() { 
        fmt.Printf("Server says: %s\n", scanner.Text()) 
    }
}