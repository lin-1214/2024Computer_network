package main
import (
	"net"
    "bufio"
    "fmt"
)

func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 
 
func main() {
    // open server connection 
    ln, _ := net.Listen("tcp", ":12000")
	conn, _ := ln.Accept() 
    defer ln.Close() 
    defer conn.Close()

    // write message to client when receive message
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)
    message, err := reader.ReadString('\n')
    check(err)

    if (message == "UPLOAD\n") {
        writer.WriteString("200 OK for UPLOAD\n")
        writer.Flush()
    }

    // receive content until EOF
    content := ""
    for {
        line, err := reader.ReadString('\n')
        check(err)
		
        if (line == "EOF\n") {
			break
		}
		content := content + line
    }

    fmt.Printf("Number of bytes received: %v\n", len(content))
}