package main

import (
	"fmt"
	"net"
	"bufio"
)

func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
}

func main() {
	conn, errc := net.Dial("tcp", "140.112.42.221:11993")
	check(errc)
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	writer.WriteString("76\n")
	writer.Flush()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() { 
        fmt.Printf("Server says: %s\n", scanner.Text()) 
    }
}