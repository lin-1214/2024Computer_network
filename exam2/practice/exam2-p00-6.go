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
	
	// write message to server
	writer := bufio.NewWriter(conn)
	writer.WriteString("UPLOAD\n")
	writer.Flush()

	// print message sent by server
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() { 
        fmt.Printf("%s\n", scanner.Text()) 
    }

	// ask for inputfile
	fmt.Printf("Input filename: ")
	input_filename := ""
	fmt.Scanf("%s", &input_filename)

	// preprocess the content
	_, err := os.Open(input_filename)
	contentBytes, err := os.ReadFile(input_filename)
	check(err)
	content := string(contentBytes)

	// write content to server
	writer.WriteString(content)
	writer.Flush()

	// send EOF
	writer.WriteString("\nEOF\n")
	writer.Flush()

	// print message sent by server
	if scanner.Scan() { 
        fmt.Printf("%s\n", scanner.Text()) 
    }
}