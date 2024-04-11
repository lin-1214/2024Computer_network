package main

import (
	"net"
	"bufio"
)

func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
}

func main() {
	conn, errc := net.Dial("tcp", "140.112.42.221:11992")
	check(errc)
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	writer.WriteString("50\n")
}