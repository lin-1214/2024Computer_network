package main

import (
	"fmt"
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
}