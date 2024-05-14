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
	writer.WriteString("PLAY\n")
	writer.Flush()

	// print message sent by server
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() { 
        fmt.Printf("%s\n", scanner.Text()) 
    }

	// ask for character
	fmt.Printf("Guess a character: ")
	ch := ""
	fmt.Scanf("%s", &ch)

	// write character to server
	writer.WriteString(fmt.Sprintf("%v\n", ch))
	writer.Flush()

	// print message sent by server
	if scanner.Scan() { 
        fmt.Printf("%s\n", scanner.Text()) 
    }
}